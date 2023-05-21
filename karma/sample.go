package karma

import (
	"math/rand"
	"time"
)

func SampleTrips(numTrips int) []TripData {
	trips := make([]TripData, 0)

	rand.Seed(time.Now().UnixNano())

	remaining := numTrips
	for remaining > 0 {
		car := generateCar()
		numTripsForCar := generateRandomInt(1) + 1

		lastStartTime := time.Now()

		for i := 0; i < numTripsForCar; i++ {

			randomHour := time.Hour * time.Duration(generateRandomInt(2)+1)
			endTime := lastStartTime.Add(-randomHour)
			startTime := generateStartTripTime(endTime)
			startDate := startTime.Format("2006-01-02 15:04:05")
			stopDate := endTime.Format("2006-01-02 15:04:05")
			lastStartTime = startTime

			tripDuration := endTime.Sub(startTime)

			startLoc := generateLocation()
			endLoc := generateLocation()

			tripDistance := calculateDistance(startLoc, endLoc)

			tripActive := 0
			if i == 0 && numTripsForCar%2 == 0 {
				tripActive = 1
				stopDate = ""
			}

			trip := TripData{
				TripID:            generateRandomString(7),
				CarID:             car.CarID,
				DriverID:          car.DriverID,
				CarNumber:         car.CarNumber,
				DeviceID:          car.DeviceID,
				TripActive:        int32(tripActive),
				StartMessageID:    generateRandomString(9),
				StartDate:         startDate,
				StartLatitude:     startLoc.Latitude,
				StartLongitude:    startLoc.Longitude,
				StartOdo:          356218.649,
				StopMessageID:     generateRandomString(9),
				StopDate:          stopDate,
				StopLatitude:      endLoc.Latitude,
				StopLongitude:     endLoc.Longitude,
				StopOdo:           356319.582,
				TripDuration:      int32(tripDuration.Minutes()),
				TripDistance:      tripDistance,
				TripDurationNight: 0,
				TripDistanceNight: 0,
			}

			trips = append(trips, trip)
		}

		remaining -= numTripsForCar
	}

	return trips
}

func SampleLocations(numLocations int) []LocationData {
	locations := make([]LocationData, 0)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numLocations; i++ {
		location := LocationData{
			MessageID: generateRandomString(9),
			CarID:     rand.Intn(100) + 1,
			CarNumber: generateRandomString(10),
			DeviceID:  generateRandomString(15),
			Extra:     generateRandomString(3),
			EDT:       generateRandomDateTime(),
			EID:       rand.Intn(3) + 1,
			Latitude:  generateRandomCoordinate(-90, 90),
			Longitude: generateRandomCoordinate(-180, 180),
			Head:      rand.Intn(360),
			Odo:       rand.Float64() * 1000000,
			Alt:       rand.Float64() * 200,
		}

		locations = append(locations, location)
	}

	return locations
}
