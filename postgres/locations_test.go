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

func TestStoreLocation(t *testing.T) {
	locations := karma.SampleLocations(1000)
	repo := postgres.Repository{Conn: conn}
	err := repo.StoreLocations(context.TODO(), locations)
	require.NoError(t, err)

	var buf bytes.Buffer
	_, err = conn.PgConn().CopyTo(context.TODO(), &buf, "COPY (SELECT json_agg(row_to_json(locations)) FROM locations) TO stdout")
	require.NoError(t, err)

	res := make([]karma.LocationData, 0, len(locations))
	err = json.NewDecoder(&buf).Decode(&res)
	require.NoError(t, err)

	assert.Equal(t, locations, res)
}
