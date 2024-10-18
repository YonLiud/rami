package database

import (
	"log"

	"rami/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("rami.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.AutoMigrate(&models.Guest{}, &models.Log{}, &models.CSO{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Connected to database")
	return nil
}
