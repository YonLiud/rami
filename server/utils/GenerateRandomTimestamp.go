package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomTimestamp() int {
	return int(time.Now().Unix()) + rand.Intn(1000000)
}
