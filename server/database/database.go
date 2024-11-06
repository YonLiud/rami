package database

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB initializes the database connection using Gorm
func InitDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// AutoMigrate your models here
		db.AutoMigrate(&models.Guest{}) // AutoMigrate the Guest model
		log.Println("Database connection established and models migrated")
	})
	return db
}

// GetDB returns the singleton instance of the database
func GetDB() *gorm.DB {
	return db
}
