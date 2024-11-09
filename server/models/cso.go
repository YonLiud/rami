package models

import (
	"gorm.io/gorm"
)

type CSO struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;not null;default:null" json:"username"`
	HashedPassword string `gorm:"not null;default:null" json:"hashedPassword"`
	Active         bool   `gorm:"default:true" json:"active"`
}
