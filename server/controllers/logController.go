package controllers

import (
	"rami/services"
)

// LogController struct to group related functions
type LogController struct {
	LogService *services.LogService
}
