package services

import (
	"rami/models"

	"gorm.io/gorm"
)

type VisitorService struct {
	DB *gorm.DB
}

func NewVisitorService(db *gorm.DB) *VisitorService {
	return &VisitorService{DB: db}
}

func (vs *VisitorService) GetAllVisitors() ([]models.Visitor, error) {
	var visitors []models.Visitor
	if err := vs.DB.Find(&visitors).Error; err != nil {
		return nil, err
	}
	return visitors, nil
}

func (vs *VisitorService) CreateVisitor(visitor *models.Visitor) error {
	if err := vs.DB.Create(visitor).Error; err != nil {
		return err
	}
	return nil
}

func (vs *VisitorService) GetVisitorByCredentialsNumber(credentialsNumber string) (models.Visitor, error) {
	var visitor models.Visitor
	if err := vs.DB.Where("credentials_number = ?", credentialsNumber).First(&visitor).Error; err != nil {
		return visitor, err
	}
	return visitor, nil
}

func (vs *VisitorService) UpdateVisitor(visitor *models.Visitor) error {
	if err := vs.DB.Save(visitor).Error; err != nil {
		return err
	}
	return nil
}
