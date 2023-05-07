package postgres_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/stretchr/testify/require"
)

type optTrips func(trips *[]karma.Trip)

func generateTrips(opts ...optTrips) []karma.Trip {
	var idCounter int32
	var trips = make([]karma.Trip, 0, 3)
	for idCounter <= 3 {
		idCounter++
		trips = append(trips, karma.Trip{
			TripID:            fmt.Sprintf("tripid-%d", idCounter),
			CarID:             idCounter,
			DriverID:          idCounter,
			CarNumber:         fmt.Sprintf("cardnumber-%d", idCounter),
			DeviceID:          fmt.Sprintf("deviceid-%d", idCounter),
			TripActive:        1,
			StartMessageID:    fmt.Sprintf("startmessageid-%d", idCounter),
			StartDate:         time.Now().Format("2006-01-0215:04:05"),
			StartLatitude:     0,
			StartLongitude:    0,
			StartOdo:          0,
			StopMessageID:     fmt.Sprintf("stopmessageid-%d", idCounter),
			StopDate:          time.Now().Format("2006-01-0215:04:05"),
			StopLatitude:      0,
			StopLongitude:     0,
			StopOdo:           0,
			TripDuration:      0,
			TripDistance:      0,
			TripDurationNight: 0,
			TripDistanceNight: 0,
		})
	}

	for _, opt := range opts {
		opt(&trips)
	}

	return trips
}

func TestStoreTrips(t *testing.T) {
	err := postgres.StoreTrips(context.TODO(), conn, generateTrips())
	require.NoError(t, err)

	f, err := os.Create("trips.csv")
	require.NoError(t, err)
	_, err = conn.PgConn().CopyTo(context.TODO(), f, "COPY trips TO stdout")
	require.NoError(t, err)
}
