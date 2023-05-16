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
	if err := run(); err != nil {
		log.Printf("failed to run: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	karmaURL := flag.String("karma-url", "http://localhost:56000", "karma api url")
	connString := flag.String("conn-string", "postgresql://localhost/vehicle-tracking?user=metabase&password=secret&sslmode=disable", "postgres db connection string")
	flag.Parse()

	ctx := context.Background()

	if err := postgres.Migrate("file://postgres/migrations", *connString); err != nil {
		return errors.Wrap(err, "could not run postgresql migrations")
	}

	conn, err := pgx.Connect(ctx, *connString)
	if err != nil {
		return errors.Wrap(err, "could not connect to postgres")
	}
	defer conn.Close(ctx)

	client := &karma.Client{URL: *karmaURL}
	repository := &postgres.Repository{Conn: conn}
	svc := syncing.NewService(client, repository)

	if err := svc.SyncTrips(ctx); err != nil {
		return errors.Wrap(err, "could not sync trips")
	}

	return nil
}
