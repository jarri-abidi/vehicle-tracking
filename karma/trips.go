package karma

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const CmdTrips = "trips"

type FetchTripsRequest struct {
	Cmd string `json:"cmd"`
}

type FetchTripsResponse struct {
	Data []TripData `json:"data"`
}

type TripData struct {
	TripID            string  `json:"trip_id"`
	CarID             int32   `json:"car_id"`     // const
	DriverID          int32   `json:"driver_id"`  // const
	CarNumber         string  `json:"car_number"` // const
	DeviceID          string  `json:"device_id"`  // const
	TripActive        int32   `json:"trip_active"`
	StartMessageID    string  `json:"start_message_id"`
	StartDate         string  `json:"start_date"`
	StartLatitude     float64 `json:"start_latitude"`
	StartLongitude    float64 `json:"start_longitude"`
	StartOdo          float64 `json:"start_odo"`
	StopMessageID     string  `json:"stop_message_id"`
	StopDate          string  `json:"stop_date"`
	StopLatitude      float64 `json:"stop_latitude"`
	StopLongitude     float64 `json:"stop_longitude"`
	StopOdo           float64 `json:"stop_odo"`
	TripDuration      int32   `json:"trip_duration"`
	TripDistance      float64 `json:"trip_distance"`
	TripDurationNight int32   `json:"trip_duration_night"`
	TripDistanceNight int32   `json:"trip_distance_night"`
}

func (c Client) FetchTrips(ctx context.Context, req FetchTripsRequest) (*FetchTripsResponse, error) {
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal request")
	}

	resp, err := http.Post(c.URL, "application/json", bytes.NewReader(buf))
	if err != nil {
		return nil, errors.Wrap(err, "could not send http request")
	}
	defer resp.Body.Close()

	var fetchTripsResponse FetchTripsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fetchTripsResponse); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}

	return &fetchTripsResponse, nil
}
