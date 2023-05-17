package main

import (
	"encoding/json"
	"net/http"

	"github.com/jarri-abidi/vehicle-tracking/karma"
	"github.com/matryer/way"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

type server struct{}

func NewMockServer() http.Handler {
	s := server{}

	router := way.NewRouter()

	router.Handle("GET", "/trips", s.handleListTrips())
	router.Handle("POST", "/locations", s.handleListLocations())

	return router
}

func (s server) handleListTrips() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trips := karma.SampleTrips(10000)

		w.Header().Set(contentTypeKey, contentTypeValue)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(trips)
	}
}

func (s server) handleListLocations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locations := karma.SampleLocations(10000)

		w.Header().Set(contentTypeKey, contentTypeValue)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(locations)
	}
}
