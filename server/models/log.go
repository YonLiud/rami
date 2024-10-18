package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Event   string `json:"action"`
	GuestID string `json:"guest"`
}
