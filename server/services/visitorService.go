package services

import (
	"rami/models"

	"gorm.io/gorm"
)

// VisitorService struct contains the database instance
type VisitorService struct {
	DB *gorm.DB
}

// NewVisitorService initializes a new VisitorService with the given database instance
func NewVisitorService(db *gorm.DB) *VisitorService {
	return &VisitorService{DB: db}
}

// GetAllVisitors retrieves all visitors from the database
func (vs *VisitorService) GetAllVisitors() ([]models.Visitor, error) {
	var visitors []models.Visitor
	if err := vs.DB.Find(&visitors).Error; err != nil {
		return nil, err
	}
	return visitors, nil
}

// CreateVisitor adds a new visitor record to the database
func (vs *VisitorService) CreateVisitor(visitor *models.Visitor) error {
	if err := vs.DB.Create(visitor).Error; err != nil {
		return err
	}
	return nil
}

// GetVisitorByCredentialsNumber retrieves a visitor by their credentials number
func (vs *VisitorService) GetVisitorByCredentialsNumber(credentialsNumber string) (models.Visitor, error) {
	var visitor models.Visitor
	if err := vs.DB.Where("credentials_number = ?", credentialsNumber).First(&visitor).Error; err != nil {
		return visitor, err
	}
	return visitor, nil
}
