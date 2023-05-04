package karma

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const cmdTrips = "trips"

type FetchTripsRequest struct {
	Cmd string `json:"cmd"`
}

type FetchTripsResponse struct {
	Data []Trip `json:"data"`
}

type Trip struct {
	TripID            string  `json:"trip_id"`
	CarID             int     `json:"car_id"`
	DriverID          int     `json:"driver_id"`
	CarNumber         string  `json:"car_number"`
	DeviceID          string  `json:"device_id"`
	TripActive        int     `json:"trip_active"`
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
	TripDuration      int     `json:"trip_duration"`
	TripDistance      float64 `json:"trip_distance"`
	TripDurationNight int     `json:"trip_duration_night"`
	TripDistanceNight int     `json:"trip_distance_night"`
}

func FetchTrips(ctx context.Context, url string, req FetchTripsRequest) (*FetchTripsResponse, error) {
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal request")
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		return nil, errors.Wrap(err, "could not send http request")
	}

	var fetchTripsResponse FetchTripsResponse
	if err := json.NewDecoder(resp.Body).Decode(&fetchTripsResponse); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}

	return &fetchTripsResponse, nil
}
