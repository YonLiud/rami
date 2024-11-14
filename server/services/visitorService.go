package services

import (
	"rami/database"
	"rami/models"
)

type VisitorService struct {
	db *database.ExcelDB
}

func NewVisitorService(db *database.ExcelDB) *VisitorService {
	return &VisitorService{db: db}
}

func (service *VisitorService) GetAllVisitors() []models.Visitor {
	return service.db.GetVisitors()
}
