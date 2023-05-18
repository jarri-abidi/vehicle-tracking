package postgres

import (
	"context"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres/gen"
)

func (r *Repository) StoreLocations(ctx context.Context, locations []karma.LocationData) error {
	queries := gen.New(r.Conn)

	args := make([]gen.InsertLocationsParams, 0, len(locations))

	for _, location := range locations {
		args = append(args, gen.InsertLocationsParams{
			MessageID: location.MessageID,
			CarID:     location.CarID,
			Carnumber: location.CarNumber,
			DeviceID:  location.DeviceID,
			Extra:     location.Extra,
			Edt:       location.EDT,
			Eid:       location.EID,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
			Head:      location.Head,
			Odo:       location.Odo,
			Alt:       location.Alt,
		})
	}

	_, err := queries.InsertLocations(ctx, args)
	return err
}
