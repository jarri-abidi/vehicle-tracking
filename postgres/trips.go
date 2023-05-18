package postgres

import (
	"context"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres/gen"
)

func (r *Repository) StoreTrips(ctx context.Context, trips []karma.TripData) error {
	queries := gen.New(r.Conn)

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
