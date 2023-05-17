package main

import (
	"context"
	"os"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v5"
	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/jarri-abidi/vehicle-tracking/syncing"
	"github.com/pkg/errors"
)

func main() {
	runtime.Start(handler)
}

func handler(ctx context.Context) error {
	karmaURL := os.Getenv("KARMA_URL")
	connString := os.Getenv("POSTGRES_URL")

	conn, err := pgx.Connect(ctx, connString)
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

	if err := svc.SyncLocations(ctx); err != nil {
		return errors.Wrap(err, "could not sync trips")
	}

	return nil
}
