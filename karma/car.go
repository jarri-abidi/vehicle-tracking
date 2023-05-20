package karma

import (
	"math/rand"
	"time"
)

type car struct {
	CarID     int32  `json:"car_id"`     // const
	DriverID  int32  `json:"driver_id"`  // const
	CarNumber string `json:"car_number"` // const
	DeviceID  string `json:"device_id"`  // const
}

func generateCar() car {
	return car{
		CarID:     generateRandomInt32(6),
		CarNumber: generateLicensePlate(),
		DriverID:  generateRandomInt32(8),
		DeviceID:  generateRandomNumbers(),
	}
}

func generateLicensePlate() string {
	rand.Seed(time.Now().UnixNano())

	// Generate one or up to two letters
	numLetters := rand.Intn(2) + 1
	letters := generateRandomLetters(numLetters)

	// Generate numbers
	numbers := generateRandomNumbers()

	// Construct the license plate
	licensePlate := letters + numbers

	return licensePlate
}

func generateRandomInt32(length int) int32 {
	rand.Seed(time.Now().UnixNano())

	min := int32(1) * int32(pow(10, length-1))
	max := int32(pow(10, length) - 1)

	return rand.Int31n(max-min+1) + min
}

func pow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
