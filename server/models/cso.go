package models

import (
	"gorm.io/gorm"
)

type CSO struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;not null" json:"username"` // Add the unique constraint for the database and keep JSON tag
	HashedPassword string `gorm:"not null" json:"hashedPassword"`       // Ensure the field is not null in the DB
	Active         bool   `gorm:"default:true" json:"active"`           // Default value for active field, also included in JSON
}
