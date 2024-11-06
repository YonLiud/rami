package database

import (
	"log"
	"rami/models"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB(dbName string) *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// Migrate models
		db.AutoMigrate(&models.Visitor{})
		db.AutoMigrate(&models.Log{})
		db.AutoMigrate(&models.CSO{})

		log.Println("Database connection established and models migrated")
	})
	return db
}

func GetDB() *gorm.DB {
	return db
}
