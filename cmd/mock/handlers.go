package main

import (
	"encoding/json"
	"net/http"

	"github.com/jarri-abidi/vehicle-tracking/karma"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

func handleListTrips(w http.ResponseWriter, r *http.Request) {
	trips := karma.SampleTrips(10000)

	w.Header().Set(contentTypeKey, contentTypeValue)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trips)
}

func handleListLocations(w http.ResponseWriter, r *http.Request) {
	locations := karma.SampleLocations(10000)

	w.Header().Set(contentTypeKey, contentTypeValue)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
