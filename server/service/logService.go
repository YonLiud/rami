package service

import (
	"rami/database"
	"rami/models"
)

func CreateLog(log *models.Log) (err error) {
	db := database.GetDB()
	if err := db.Create(log).Error; err != nil {
		return err
	}
	return nil
}

func GetAllLogs() (logs []models.Log, err error) {
	db := database.GetDB()
	if err := db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func GetLogBySerial(serial int) (log models.Log, err error) {
	db := database.GetDB()
	if err := db.First(&log, serial).Error; err != nil {
		return log, err
	}
	return log, nil
}
