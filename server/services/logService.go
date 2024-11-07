package services

import (
	"rami/models"

	"gorm.io/gorm"
)

// LogService provides methods to work with logs in the database
type LogService struct {
	DB *gorm.DB
}

// NewLogService creates a new instance of LogService
func NewLogService(db *gorm.DB) *LogService {
	return &LogService{DB: db}
}

// CreateLog saves a new log entry to the database
func (s *LogService) CreateLog(log *models.Log) error {
	if err := s.DB.Create(log).Error; err != nil {
		return err
	}
	return nil
}

// GetAllLogs retrieves all logs from the database
func (s *LogService) GetAllLogs() ([]models.Log, error) {
	var logs []models.Log
	if err := s.DB.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

// GetLogsBySerial retrieves logs based on a specific serial from the database
func (s *LogService) GetLogsBySerial(serial string) ([]models.Log, error) {
	var logs []models.Log
	if err := s.DB.Where("serial = ?", serial).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
