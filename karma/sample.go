package karma

import (
	"math/rand"
	"time"
)

func SampleTrips(numTrips int32) []TripData {
	trips := make([]TripData, 0)

	rand.Seed(time.Now().UnixNano())

	remaining := numTrips
	for remaining > 0 {
		car := generateCar()
		numTripsForCar := generateRandomInt32(1) + 1

		// generate trips for car and append to trips
		for i := 0; i < int(numTripsForCar); i++ {
			tripActive := 0
			if i == int(numTripsForCar)-1 && numTripsForCar%2 == 0 {
				tripActive = 1
			}
			trip := TripData{
				TripID:            generateRandomString(7),
				CarID:             car.CarID,
				DriverID:          car.DriverID,
				CarNumber:         car.CarNumber,
				DeviceID:          car.DeviceID,
				TripActive:        int32(tripActive),
				StartMessageID:    generateRandomString(9),
				StartDate:         generateRandomDateTime(),
				StartLatitude:     31.958073,
				StartLongitude:    34.847893,
				StartOdo:          356218.649,
				StopMessageID:     generateRandomString(9),
				StopDate:          generateRandomDateTime(),
				StopLatitude:      31.830221,
				StopLongitude:     35.236596,
				StopOdo:           356319.582,
				TripDuration:      rand.Int31n(200),
				TripDistance:      rand.Float64() * 200,
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
