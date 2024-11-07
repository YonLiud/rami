package utils

import (
	"rami/database"

	"gorm.io/gorm"
)

func InitiateTestDB() *gorm.DB {
	db := database.InitDB(":memory:")

	db.Exec("DELETE FROM visitors;")
	db.Exec("DELETE FROM logs;")
	db.Exec("DELETE FROM csos;")

	return db
}
