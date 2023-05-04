package karma

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchTrips(t *testing.T) {
	f, err := os.Open("fetch_trips_sample.json")
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req FetchTripsRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, cmdTrips, req.Cmd)

		_, err := io.Copy(w, f)
		require.NoError(t, err)
	}))

	_, err = FetchTrips(context.TODO(), srv.URL, FetchTripsRequest{Cmd: cmdTrips})
	require.NoError(t, err)
}
