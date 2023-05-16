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

func TestFetchLocation(t *testing.T) {
	f, err := os.Open("testdata/location_sample.json")
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req karma.FetchLocationRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, karma.CmdLocation, req.Cmd)

		_, err := io.Copy(w, f)
		require.NoError(t, err)
	}))

	c := karma.Client{URL: srv.URL}
	_, err = c.FetchLocation(context.TODO(), karma.FetchLocationRequest{Cmd: karma.CmdLocation})
	require.NoError(t, err)
}
