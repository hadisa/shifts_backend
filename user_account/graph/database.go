package graph

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseParameters struct {
	hostname string
	port     string
	name     string
	timeout  time.Duration
}

// DatabaseCredentials is a set of dynamic credentials retrieved from Vault
type DatabaseCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Database struct {
	connection      *gorm.DB
	connectionMutex sync.Mutex
	parameters      DatabaseParameters
}

// NewDatabase establishes a database connection with the given Vault credentials
func NewDatabase(ctx context.Context, parameters DatabaseParameters, credentials DatabaseCredentials) (*Database, error) {
	database := &Database{
		connection:      nil,
		connectionMutex: sync.Mutex{},
		parameters:      parameters,
	}

	// establish the first connection
	if err := database.Reconnect(ctx, credentials); err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return database, nil
}

// Reconnect will be called periodically to refresh the database connection
// since the dynamic credentials expire after some time, it will:
//  1. construct a connection string using the given credentials
//  2. establish a database connection
//  3. close & replace the existing connection with the new one behind a mutex
func (db *Database) Reconnect(ctx context.Context, credentials DatabaseCredentials) error {
	ctx, cancelContextFunc := context.WithTimeout(ctx, db.parameters.timeout)
	defer cancelContextFunc()

	log.Printf(
		"connecting to %q database @ %s:%s with username %q",
		db.parameters.name,
		db.parameters.hostname,
		db.parameters.port,
		credentials.Username,
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		db.parameters.hostname,
		db.parameters.port,
		db.parameters.name,
		credentials.Username,
		credentials.Password,
	)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return fmt.Errorf("unable to open database connection: %w", err)
	}

	// wait until the database is ready or timeout expires
	for {
		sqlDB, err := connection.DB()
		// Ping
		err = sqlDB.Ping()
		if err == nil {
			break
		}

		select {
		case <-time.After(500 * time.Millisecond):
			continue
		case <-ctx.Done():
			return fmt.Errorf("failed to successfully ping database before context timeout: %w", err)
		}
	}

	db.closeReplaceConnection(connection)
	db.GetConnection()
	log.Printf("connecting to %q database: success!", db.parameters.name)

	return nil
}

// close & replace the existing connection with the new one behind a mutex
func (db *Database) closeReplaceConnection(new *gorm.DB) {
	/* */ db.connectionMutex.Lock()
	defer db.connectionMutex.Unlock()

	if db.connection != nil {
		sqlDB, err := db.connection.DB()
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)
		}
		// close the connection
		sqlDB.Close()
	}

	// replace the connection with the new one
	db.connection = new

}

func (db *Database) Close() error {
	/* */ db.connectionMutex.Lock()
	defer db.connectionMutex.Unlock()

	if db.connection != nil {
		sqlDB, err := db.connection.DB()
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)
		}
		// close the connection
		return sqlDB.Close()
	}

	return nil
}

// share the database connection with other packages by using a global variable
var dbConn gorm.DB

func (db *Database) GetConnection() {
	/* */ db.connectionMutex.Lock()
	defer db.connectionMutex.Unlock()
	dbConn = *db.connection

}

// call the function to retrieve the database connection
func GetOpenConnection() *gorm.DB {
	return &dbConn
}

// Use models to create tables
func CreateTables() {
	db := GetOpenConnection()
	db.AutoMigrate()
}
