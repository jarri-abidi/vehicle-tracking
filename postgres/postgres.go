package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres/gen"
	"github.com/pkg/errors"
)

const dbName = "karma"

func pgConnectionString(host, port, user, pgPassword string, options ...string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", user, pgPassword, host, port, dbName, strings.Join(options, "&"))
}

func Migrate(migrationURL, host, port, user, password string, options ...string) error {
	db, err := sql.Open("postgres", pgConnectionString(host, port, user, password, options...))
	if err != nil {
		return errors.Wrap(err, "could not open postgres connection")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()
	driver, err := postgres.WithInstance(db, &postgres.Config{MultiStatementEnabled: true})
	if err != nil {
		return errors.Wrap(err, "could not create postgres driver instance")
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationURL, dbName, driver)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to run migrate up")
	}

	return nil
}

func StoreTrips(ctx context.Context, conn *pgx.Conn, trips []karma.Trip) error {
	queries := gen.New(conn)

	args := make([]gen.InsertTripsParams, 0, len(trips))
	for _, trip := range trips {
		args = append(args, gen.InsertTripsParams{
			TripID:            trip.TripID,
			CarID:             trip.CarID,
			DriverID:          trip.DriverID,
			CarNumber:         trip.CarNumber,
			DeviceID:          trip.DeviceID,
			TripActive:        trip.TripActive,
			StartMessageID:    trip.StartMessageID,
			StartDate:         trip.StartDate,
			StartLatitude:     trip.StartLatitude,
			StartLongitude:    trip.StartLongitude,
			StartOdo:          trip.StartOdo,
			StopMessageID:     trip.StopMessageID,
			StopDate:          trip.StopDate,
			StopLatitude:      trip.StopLatitude,
			StopLongitude:     trip.StopLongitude,
			StopOdo:           trip.StopOdo,
			TripDuration:      trip.TripDuration,
			TripDistance:      trip.TripDistance,
			TripDurationNight: trip.TripDurationNight,
			TripDistanceNight: trip.TripDistanceNight,
		})
	}

	_, err := queries.InsertTrips(ctx, args)
	return err
}
