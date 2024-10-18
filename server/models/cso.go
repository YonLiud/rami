package models

import (
	"gorm.io/gorm"
)

type CSO struct {
	gorm.Model
	UserName       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}
