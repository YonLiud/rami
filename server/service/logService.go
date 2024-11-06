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

func GetLogsBySerial(serial string) (logs []models.Log, err error) {
	db := database.GetDB()
	if err := db.Where("serial = ?", serial).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
