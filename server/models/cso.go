package models

import (
	"gorm.io/gorm"
)

type CSO struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;not null" json:"username"`
	HashedPassword string `gorm:"not null" json:"hashedPassword"`
	Active         bool   `gorm:"default:true" json:"active"`
}
