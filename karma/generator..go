package karma

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var (
	letterRunes   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitRunes    = []rune("0123456789")
	earthRadiusKm = 6371.0
)

func generateRandomLetters(numLetters int) string {
	var sb strings.Builder
	for i := 0; i < numLetters; i++ {
		if i == 1 {
			sb.WriteString("A")
		} else {
			letter := string(letterRunes[rand.Intn(len(letterRunes))])
			sb.WriteString(letter)
		}
	}
	return sb.String()
}

func generateRandomNumbers() string {
	numDigits := rand.Intn(5) + 1
	var sb strings.Builder
	for i := 0; i < numDigits; i++ {
		digit := string(digitRunes[rand.Intn(len(digitRunes))])
		sb.WriteString(digit)
	}
	return sb.String()
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

func generateStartTripTime(endTime time.Time) time.Time {
	randomMin := generateRandomInt(2) + 1
	randomDelta := time.Minute * time.Duration(randomMin)

	startTime := endTime.Add(-randomDelta)

	return startTime
}

func generateRandomCoordinate(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

type location struct {
	Latitude  float64
	Longitude float64
}

func generateLocation() location {
	for _, loc := range locations {
		return loc
	}
	return location{}
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func calculateDistance(loc1, loc2 location) float64 {
	// Convert latitude and longitude to radians
	lat1Rad := degreesToRadians(loc1.Latitude)
	lon1Rad := degreesToRadians(loc1.Longitude)
	lat2Rad := degreesToRadians(loc2.Latitude)
	lon2Rad := degreesToRadians(loc2.Longitude)

	// Calculate the differences between the coordinates
	latDiff := lat2Rad - lat1Rad
	lonDiff := lon2Rad - lon1Rad

	// Calculate the Haversine formula
	a := math.Pow(math.Sin(latDiff/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(lonDiff/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadiusKm * c

	return distance
}

var locations = map[string]location{
	"Burj Khalifa": {
		Latitude:  25.1972,
		Longitude: 55.2744,
	},
	"The Dubai Mall": {
		Latitude:  25.1972,
		Longitude: 55.2791,
	},
	"Dubai Marina": {
		Latitude:  25.0775,
		Longitude: 55.1406,
	},
	"Palm Jumeirah": {
		Latitude:  25.1126,
		Longitude: 55.1388,
	},
	"Dubai Creek": {
		Latitude:  25.2667,
		Longitude: 55.3000,
	},
	"Jumeirah Beach Residence": {
		Latitude:  25.0760,
		Longitude: 55.1374,
	},
	"Dubai Frame": {
		Latitude:  25.1959,
		Longitude: 55.2750,
	},
	"Dubai Opera": {
		Latitude:  25.1949,
		Longitude: 55.2756,
	},
	"Mall of the Emirates": {
		Latitude:  25.1186,
		Longitude: 55.2008,
	},
	"Dubai Miracle Garden": {
		Latitude:  25.0608,
		Longitude: 55.2354,
	},
	"Ibn Battuta Mall": {
		Latitude:  25.0442,
		Longitude: 55.1198,
	},
	"Dubai Parks and Resorts": {
		Latitude:  25.0811,
		Longitude: 55.1179,
	},
	"Dubai Festival City": {
		Latitude:  25.2270,
		Longitude: 55.3582,
	},
	"Dubai Creek Park": {
		Latitude:  25.2277,
		Longitude: 55.3300,
	},
	"Deira City Centre": {
		Latitude:  25.2647,
		Longitude: 55.3328,
	},
	"Dubai Sports City": {
		Latitude:  25.0324,
		Longitude: 55.1900,
	},
	"Dubai Marina Yacht Club": {
		Latitude:  25.0775,
		Longitude: 55.1383,
	},
	"Kite Beach": {
		Latitude:  25.1653,
		Longitude: 55.2175,
	},
	"Dubai Butterfly Garden": {
		Latitude:  25.1070,
		Longitude: 55.1734,
	},
	"La Mer": {
		Latitude:  25.2326,
		Longitude: 55.2560,
	},
	"Dubai Investment Park": {
		Latitude:  24.9858,
		Longitude: 55.1934,
	},
}
