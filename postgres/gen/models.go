// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package gen

import ()

type Trip struct {
	TripID            string
	CarID             int32
	DriverID          int32
	CarNumber         string
	DeviceID          string
	TripActive        int32
	StartMessageID    string
	StartDate         string
	StartLatitude     float64
	StartLongitude    float64
	StartOdo          float64
	StopMessageID     string
	StopDate          string
	StopLatitude      float64
	StopLongitude     float64
	StopOdo           float64
	TripDuration      int32
	TripDistance      float64
	TripDurationNight int32
	TripDistanceNight int32
}
