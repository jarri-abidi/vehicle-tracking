package postgres_test

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreTrips(t *testing.T) {
	trips := karma.SampleTrips(1000)
	repo := postgres.Repository{Conn: conn}
	err := repo.StoreTrips(context.TODO(), trips)
	require.NoError(t, err)

	var buf bytes.Buffer
	_, err = conn.PgConn().CopyTo(context.TODO(), &buf, "COPY (SELECT json_agg(row_to_json(trips)) FROM trips) TO stdout")
	require.NoError(t, err)

	res := make([]karma.TripData, 0, len(trips))
	err = json.NewDecoder(&buf).Decode(&res)
	require.NoError(t, err)

	assert.Equal(t, trips, res)
}

// goos: linux
// goarch: amd64
// pkg: github.com/jarri-abidi/vehicle-tracking/postgres
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkStoreTrips-4   	       5	 883662450 ns/op	38482932 B/op	  619525 allocs/op
func BenchmarkStoreTrips(b *testing.B) {
	trips := karma.SampleTrips(50000)
	repo := postgres.Repository{Conn: conn}

	for i := 0; i < b.N; i++ {
		require.NoError(b, repo.StoreTrips(context.TODO(), trips))
	}
}
