package graph

import (
	"account_user/graph/generated"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
)

const defaultPort = "8080"

// const defaultPort = "8110"

func run(ctx context.Context) error {
	ctx, cancelContextFunc := context.WithCancel(ctx)
	defer cancelContextFunc()

	// vault
	// vault, authToken, err := NewVaultAppRoleClient(
	// 	ctx,
	// 	VaultParameters{
	// 		address:                 os.Getenv("VAULT_ADDRESS"),
	// 		approleRoleID:           os.Getenv("VAULT_APPROLE_ROLE_ID"),
	// 		approleSecretID:         os.Getenv("VAULT_APPROLE_SECRET_ID"),
	// 		apiKeyPath:              os.Getenv("VAULT_API_KEY_PATH"),
	// 		apiKeyMountPath:         os.Getenv("VAULT_API_KEY_MOUNT_PATH"),
	// 		apiKeyField:             os.Getenv("VAULT_API_KEY_FIELD"),
	// 		databaseCredentialsPath: os.Getenv("VAULT_DATABASE_CREDS_PATH"),
	// 	},
	// )
	// if err != nil {
	// 	sentry.CaptureException(err)
	// 	defer sentry.Flush(2 * time.Second)

	// 	return fmt.Errorf("unable to initialize vault connection : %w", err)
	// }

	// database
	// databaseCredentials, databaseCredentialsLease, err := vault.GetDatabaseCredentials(ctx)
	// if err != nil {

	// 	sentry.CaptureException(err)
	// 	defer sentry.Flush(2 * time.Second)

	// 	return fmt.Errorf("unable to retrieve database credentials from vault: %w", err)
	// }

	timeOut, err := strconv.Atoi(os.Getenv("DATABASE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("unable to convert database timeout to int: %w", err)
	}
	database, err := NewDatabase(
		ctx,
		DatabaseParameters{
			hostname: os.Getenv("DATABASE_HOST"),
			port:     os.Getenv("DATABASE_PORT"),
			name:     os.Getenv("DATABASE_NAME"),
			timeout:  time.Duration(timeOut) * time.Second,
		},
		DatabaseCredentials{
			Username: "postgres",
			Password: "12345",
		},
	)
	if err != nil {
		// sentry.CaptureException(err)
		// defer sentry.Flush(2 * time.Second)

		return fmt.Errorf("unable to connect to database : %w", err)
	}

	defer func() {
		_ = database.Close()
	}()

	// start the lease-renewal goroutine & wait for it to finish on exit
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	vault.PeriodicallyRenewLeases(ctx, authToken, databaseCredentialsLease, database.Reconnect)
	// 	wg.Done()
	// }()
	// defer func() {
	// 	cancelContextFunc()
	// 	wg.Wait()
	// }()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: GetOpenConnection()}}))

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	CreateTables()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// endless.ListenAndServe(":"+port, mux) <--- this should be used in production
	log.Fatal(http.ListenAndServe(":"+port, mux))

	return nil
}

func Start() {
	if err := run(context.Background()); err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		log.Fatalf("error: %v", err)
	}
}
