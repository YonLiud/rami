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

func (s *LogService) CreateLogHelper(event string, serial string) error {
	log := &models.Log{
		Event:     event,
		Serial:    serial,
		Timestamp: models.GetCurrentTimestamp(),
	}

	if err := s.DB.Create(log).Error; err != nil {
		return err
	}
	return nil
}

func (s *LogService) CreateLog(log *models.Log) error {
	if err := s.DB.Create(log).Error; err != nil {
		return err
	}
	return nil
}

func (s *LogService) GetAllLogs() ([]models.Log, error) {
	var logs []models.Log
	if err := s.DB.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (s *LogService) GetLogsBySerial(serial string) ([]models.Log, error) {
	var logs []models.Log
	if err := s.DB.Where("serial = ?", serial).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
