package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Event     string `json:"event"`
	Serial    string `json:"serial"` // username / personal number associated with the event
	Timestamp int    `json:"timestamp"`
}
