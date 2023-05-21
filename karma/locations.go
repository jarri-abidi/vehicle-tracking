package karma

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const CmdLocation = "location"

type FetchLocationRequest struct {
	Cmd string `json:"cmd"`
}

type FetchLocationResponse struct {
	Data []LocationData `json:"data"`
}

type LocationData struct {
	MessageID string  `json:"message_id"`
	CarID     int     `json:"car_id"`    // const
	CarNumber string  `json:"carnumber"` // const
	DeviceID  string  `json:"device_id"` // const
	Extra     string  `json:"extra"`
	EDT       string  `json:"edt"`
	EID       int     `json:"eid"` // const
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Head      int32   `json:"head"`
	Odo       float64 `json:"odo"`
	Alt       float64 `json:"alt"`
}

func (c Client) FetchLocation(ctx context.Context, req FetchLocationRequest) (*FetchLocationResponse, error) {
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal request")
	}

	resp, err := http.Post(c.URL, "application/json", bytes.NewReader(buf))
	if err != nil {
		return nil, errors.Wrap(err, "could not send http request")
	}
	defer resp.Body.Close()

	var fetchLocationResponse FetchLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&fetchLocationResponse); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}

	return &fetchLocationResponse, nil
}
