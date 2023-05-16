package postgres

import (
	"context"

	"github.com/jarri-abidi/vehicle-tracking/karma"
)

func (r *Repository) StoreLocations(ctx context.Context, locations []karma.LocationData) error {
	return nil
}
