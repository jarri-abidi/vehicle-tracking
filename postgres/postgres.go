package postgres

import (
	"context"
	"database/sql"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres/gen"
)

func StoreTrips(ctx context.Context, db *sql.DB, trips []karma.Trip) error {
	queries := gen.New(db)

	return queries.InsertTrips(ctx, gen.InsertTripsParams{})
}
