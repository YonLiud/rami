package models

import (
	"time"

	"gorm.io/gorm"
)

func GetCurrentTimestamp() int {
	return int(time.Now().Unix())
}

type Log struct {
	gorm.Model
	Event     string `json:"event"`
	Serial    string `json:"serial"` // username / personal number associated with the event's subject
	Timestamp int    `json:"timestamp"`
}
