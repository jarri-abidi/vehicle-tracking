package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/jarri-abidi/vehicle-tracking/syncing"
	"github.com/pkg/errors"
)

func main() {
	karmaURL := flag.String("karma-url", "http://localhost:56000", "karma api url")
	pgConnString := flag.String("conn-string", "postgresql://localhost/vehicle-tracking?user=metabase&password=secret&sslmode=disable", "postgres db connection string")
	flag.Parse()

	if err := run(*karmaURL, *pgConnString); err != nil {
		log.Printf("failed to run: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(karmaURL, pgConnString string) error {
	if err := postgres.Migrate("file://postgres/migrations", pgConnString); err != nil {
		return errors.Wrap(err, "could not run postgresql migrations")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, pgConnString)
	if err != nil {
		return errors.Wrap(err, "could not connect to postgres")
	}
	defer conn.Close(ctx)

	client := &karma.Client{URL: karmaURL}
	repository := &postgres.Repository{Conn: conn}
	svc := syncing.NewService(client, repository)

	if err := svc.SyncTrips(ctx); err != nil {
		return errors.Wrap(err, "could not sync trips")
	}

	return nil
}
