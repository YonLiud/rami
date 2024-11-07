package services

import (
	"math/rand"
	"rami/database"
	"strings"
	"time"

	"gorm.io/gorm"
)

func InitiateTestDB() *gorm.DB {
	db := database.InitDB(":memory:")

	db.Exec("DELETE FROM visitors;")
	db.Exec("DELETE FROM logs;")
	db.Exec("DELETE FROM csos;")

	return db
}

func generateRandomTimestamp() int {
	return int(time.Now().Unix()) + rand.Intn(1000000)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(charset[rand.Intn(len(charset))])
	}
	return result.String()
}

func generateRandomChoice(options []string) string {
	return options[rand.Intn(len(options))] // Select a random index
}

func generateRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
