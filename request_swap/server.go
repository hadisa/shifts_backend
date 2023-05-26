package main

//go:generate go run github.com/99designs/gqlgen generate
import (
	"log"
	"os"
	"request_swaps/graph"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

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

	graph.Start()
}
