package karma

import (
	"math/rand"
	"strings"
	"time"
)

var (
	letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitRunes  = []rune("0123456789")
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

func generateRandomCoordinate(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
