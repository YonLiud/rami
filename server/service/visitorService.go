package service

import (
	"rami/database"
	"rami/models"
)

func GetAllVisitors() (visitors []models.Visitor, err error) {
	db := database.GetDB()
	if err := db.Find(&visitors).Error; err != nil {
		return nil, err
	}
	return visitors, nil
}

func CreateVisitor(Visitor *models.Visitor) (err error) {
	db := database.GetDB()
	if err := db.Create(Visitor).Error; err != nil {
		return err
	}
	return nil
}
