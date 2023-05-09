package postgres_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type optTrips func(trips *[]karma.Trip)

func generateTrips(tripsCount int32, opts ...optTrips) []karma.Trip {
	var idCounter int32
	var trips = make([]karma.Trip, 0, tripsCount)
	for idCounter <= tripsCount {
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
	trips := generateTrips(1000)
	err := postgres.StoreTrips(context.TODO(), conn, trips)
	require.NoError(t, err)

	var buf bytes.Buffer
	_, err = conn.PgConn().CopyTo(context.TODO(), &buf, "COPY (SELECT json_agg(row_to_json(trips)) FROM trips) TO stdout")
	require.NoError(t, err)

	gotTrips := make([]karma.Trip, 0, len(trips))
	err = json.NewDecoder(&buf).Decode(&gotTrips)
	require.NoError(t, err)

	assert.Equal(t, trips, gotTrips)
}

func BenchmarkStoreTrips(b *testing.B) {
	trips := generateTrips(50000)

	for i := 0; i < b.N; i++ {
		require.NoError(b, postgres.StoreTrips(context.TODO(), conn, trips))
	}
}
