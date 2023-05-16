package syncing

import (
	"context"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/pkg/errors"
)

type Service struct {
	client     *karma.Client
	repository *postgres.Repository
}

func NewService(client *karma.Client, repository *postgres.Repository) *Service {
	return &Service{client: client, repository: repository}
}

func (s *Service) SyncTrips(ctx context.Context) error {
	resp, err := s.client.FetchTrips(ctx, karma.FetchTripsRequest{Cmd: karma.CmdTrips})
	if err != nil {
		return errors.Wrap(err, "could not fetch trips")
	}

	if err := s.repository.StoreTrips(ctx, resp.Data); err != nil {
		return errors.Wrap(err, "could not store trips")
	}

	return nil
}
