package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"account_user/graph"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	ory "github.com/ory/client-go"

	"github.com/joho/godotenv"
)

type App struct {
	ory *ory.APIClient
}

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DNS"),
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// configuration := ory.NewConfiguration()
	// configuration.Servers = []ory.ServerConfiguration{
	// 	{
	// 		URL: os.Getenv("ORY_KRATOS_HOST"), // Kratos Admin API
	// 	},
	// }

	// app := &App{
	// 	ory: ory.NewAPIClient(configuration),
	// }

	// Run App with the Kratos Middleware to check the session
	// mux.Handle("/", app.sessionMiddleware(playground.Handler("GraphQL playground", "/query")))

	graph.Start()

}
