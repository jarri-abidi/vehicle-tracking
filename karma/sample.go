package karma

import (
	"math/rand"
	"time"
)

func SampleTrips(numTrips int) []TripData {
	trips := make([]TripData, 0)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTrips; i++ {
		trip := TripData{
			TripID:            generateRandomString(7),
			CarID:             12,
			DriverID:          33,
			CarNumber:         "My Car",
			DeviceID:          generateRandomString(15),
			TripActive:        0,
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

	return trips
}

func SampleLocations(numLocations int) []LocationData {
	locations := make([]LocationData, 0)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numLocations; i++ {
		location := LocationData{
			MessageID: generateRandomString(9),
			CarID:     rand.Int31(),
			CarNumber: generateRandomString(10),
			DeviceID:  generateRandomString(15),
			Extra:     generateRandomString(3),
			EDT:       generateRandomDateTime(),
			EID:       rand.Int31(),
			Latitude:  generateRandomCoordinate(-90, 90),
			Longitude: generateRandomCoordinate(-180, 180),
			Head:      rand.Int31(),
			Odo:       rand.Float64() * 1000000,
			Alt:       rand.Float64() * 200,
		}

		locations = append(locations, location)
	}

	return locations
}

func generateRandomString(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charSetLength := len(charSet)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = charSet[rand.Intn(charSetLength)]
	}

	return string(result)
}

func generateRandomDateTime() string {
	start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)

	delta := end.Sub(start)
	randomDelta := time.Duration(rand.Int63n(int64(delta)))

	randomTime := start.Add(randomDelta)

	return randomTime.Format("2006-01-02 15:04:05")
}

func generateRandomCoordinate(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
