package karma_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchTrips(t *testing.T) {
	f, err := os.Open("trips_sample.json")
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req karma.FetchTripsRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, karma.CmdTrips, req.Cmd)

		_, err := io.Copy(w, f)
		require.NoError(t, err)
	}))

	c := karma.Client{URL: srv.URL}
	_, err = c.FetchTrips(context.TODO(), karma.FetchTripsRequest{Cmd: karma.CmdTrips})
	require.NoError(t, err)
}
