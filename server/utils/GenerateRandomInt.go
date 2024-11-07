package utils

import "math/rand"

func GenerateRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
