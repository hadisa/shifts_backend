package graph

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	vault "github.com/hashicorp/vault/api"
)

func (v *Vault) PeriodicallyRenewLeases(
	ctx context.Context,
	authToken *vault.Secret,
	databaseCredentialsLease *vault.Secret,
	databaseReconnectFunc func(ctx context.Context, credentials DatabaseCredentials) error,
) {
	/* */ log.Println("renew / recreate secrets loop: begin")
	defer log.Println("renew / recreate secrets loop: end")

	currentAuthToken := authToken
	currentDatabaseCredentialsLease := databaseCredentialsLease

	for {
		renewed, err := v.renewLeases(ctx, currentAuthToken, currentDatabaseCredentialsLease)
		if err != nil {
			log.Fatalf("renew error: %v", err) // simplified error handling
		}

		if renewed&exitRequested != 0 {
			return
		}

		if renewed&expiringAuthToken != 0 {
			log.Printf("auth token: can no longer be renewed; will log in again")

			authToken, err := v.login(ctx)
			if err != nil {
				sentry.CaptureException(err)
				defer sentry.Flush(2 * time.Second)
				log.Fatalf("login authentication error: %v", err) // simplified error handling
			}

			currentAuthToken = authToken
		}

		if renewed&expiringDatabaseCredentialsLease != 0 {
			log.Printf("database credentials: can no longer be renewed; will fetch new credentials & reconnect")

			databaseCredentials, databaseCredentialsLease, err := v.GetDatabaseCredentials(ctx)
			if err != nil {
				sentry.CaptureException(err)
				defer sentry.Flush(2 * time.Second)
				log.Fatalf("database credentials error: %v", err) // simplified error handling
			}

			if err := databaseReconnectFunc(ctx, databaseCredentials); err != nil {
				sentry.CaptureException(err)
				defer sentry.Flush(2 * time.Second)
				log.Fatalf("database connection error: %v", err) // simplified error handling
			}

			currentDatabaseCredentialsLease = databaseCredentialsLease
		}
	}
}

// renewResult is a bitmask which could contain one or more of the values below
type renewResult uint8

const (
	renewError renewResult = 1 << iota
	exitRequested
	expiringAuthToken                // will be revoked soon
	expiringDatabaseCredentialsLease // will be revoked soon
)

// renewLeases is a blocking helper function that uses LifetimeWatcher
// instances to periodically renew the given secrets when they are close to
// their 'token_ttl' expiration times until one of the secrets is close to its
// 'token_max_ttl' lease expiration time.
func (v *Vault) renewLeases(ctx context.Context, authToken, databaseCredentialsLease *vault.Secret) (renewResult, error) {
	/* */ log.Println("renew cycle: begin")
	defer log.Println("renew cycle: end")

	// auth token
	authTokenWatcher, err := v.client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: authToken,
	})
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)
		return renewError, fmt.Errorf("unable to initialize auth token lifetime watcher: %w", err)
	}

	go authTokenWatcher.Start()
	defer authTokenWatcher.Stop()

	// database credentials
	databaseCredentialsWatcher, err := v.client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: databaseCredentialsLease,
	})
	if err != nil {
		return renewError, fmt.Errorf("unable to initialize database credentials lifetime watcher: %w", err)
	}

	go databaseCredentialsWatcher.Start()
	defer databaseCredentialsWatcher.Stop()

	// monitor events from both watchers
	for {
		select {
		case <-ctx.Done():
			return exitRequested, nil

		// DoneCh will return if renewal fails, or if the remaining lease
		// duration is under a built-in threshold and either renewing is not
		// extending it or renewing is disabled.  In both cases, the caller
		// should attempt a re-read of the secret. Clients should check the
		// return value of the channel to see if renewal was successful.
		case err := <-authTokenWatcher.DoneCh():
			// Leases created by a token get revoked when the token is revoked.
			return expiringAuthToken | expiringDatabaseCredentialsLease, err

		case err := <-databaseCredentialsWatcher.DoneCh():
			return expiringDatabaseCredentialsLease, err

		// RenewCh is a channel that receives a message when a successful
		// renewal takes place and includes metadata about the renewal.
		case info := <-authTokenWatcher.RenewCh():
			log.Printf("auth token: successfully renewed; remaining duration: %ds", info.Secret.Auth.LeaseDuration)

		case info := <-databaseCredentialsWatcher.RenewCh():
			log.Printf("database credentials: successfully renewed; remaining lease duration: %ds", info.Secret.LeaseDuration)
		}
	}
}
