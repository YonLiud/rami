package services

import (
	"errors"
	"log"
	"rami/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CSOService struct {
	DB *gorm.DB
}

// NewCSOService creates a new instance of CSOService
func NewCSOService(db *gorm.DB) *CSOService {
	return &CSOService{DB: db}
}

// HashPassword hashes a plain password for secure storage
func (s *CSOService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain password
func (s *CSOService) ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// DeactivateCSO sets the active status of a CSO to false
func (s *CSOService) DeactivateCSO(username string) error {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		return err
	}
	cso.Active = false
	return s.DB.Save(&cso).Error
}

// ActivateCSO sets the active status of a CSO to true
func (s *CSOService) ActivateCSO(username string) error {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		return err
	}
	cso.Active = true
	return s.DB.Save(&cso).Error
}

// CreateCSO creates a new CSO record in the database
func (s *CSOService) CreateCSO(username, password string) error {
	var existingCSO models.CSO
	err := s.DB.Where("username = ?", username).First(&existingCSO).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		return models.ErrUsernameExists
	}

	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return err
	}

	cso := models.CSO{
		Username:       username,
		HashedPassword: hashedPassword,
		Active:         true,
	}
	return s.DB.Create(&cso).Error
}

// GetCSOByUsername retrieves a CSO by username
func (s *CSOService) GetCSOByUsername(username string) (models.CSO, error) {
	var cso models.CSO
	err := s.DB.Where("username = ?", username).First(&cso).Error
	if err != nil {
		return cso, err
	}
	return cso, nil
}

// GetAllActiveCSOs retrieves all active CSOs
func (s *CSOService) GetAllActiveCSOs() ([]models.CSO, error) {
	var csos []models.CSO
	err := s.DB.Where("active = ?", true).Find(&csos).Error
	return csos, err
}

// AuthenticateCSO verifies the CSO's credentials
func (s *CSOService) AuthenticateCSO(username, password string) (bool, error) {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, models.ErrInvalidCredentials
		}
		return false, err
	}

	if !cso.Active {
		return false, models.ErrCSOInactive
	}

	if s.ComparePasswords(cso.HashedPassword, password) {
		return true, nil
	}

	return false, models.ErrInvalidCredentials
}
