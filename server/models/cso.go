package models

import (
	"gorm.io/gorm"
)

type CSO struct {
	gorm.Model
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

