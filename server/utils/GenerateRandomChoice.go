package utils

import "math/rand"

func GenerateRandomChoice(options []string) string {
	return options[rand.Intn(len(options))] // Select a random index
}
