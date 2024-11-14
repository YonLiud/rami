package services

import (
	"rami/models"

	"gorm.io/gorm"
)

type LogService struct {
	DB *gorm.DB
}

func NewLogService(db *gorm.DB) *LogService {
	return &LogService{DB: db}
}

func (service *LogService) GetAllLogs() []models.Log {
	var logs []models.Log
	service.DB.Find(&logs)
	return logs
}
