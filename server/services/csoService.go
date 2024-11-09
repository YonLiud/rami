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

func NewCSOService(db *gorm.DB) *CSOService {
	return &CSOService{DB: db}
}

func (s *CSOService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *CSOService) ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

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

func (s *CSOService) GetCSOByUsername(username string) (models.CSO, error) {
	var cso models.CSO
	err := s.DB.Where("username = ?", username).First(&cso).Error
	if err != nil {
		return cso, err
	}
	return cso, nil
}

func (s *CSOService) GetAllActiveCSOs() ([]models.CSO, error) {
	var csos []models.CSO
	err := s.DB.Where("active = ?", true).Find(&csos).Error
	return csos, err
}

func (s *CSOService) AuthenticateCSO(username, password string) (string, error) {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrInvalidCredentials
		}
		return "", err
	}

	if !cso.Active {
		return "", models.ErrInactiveCSO
	}

	if !s.ComparePasswords(cso.HashedPassword, password) {
		return "", models.ErrInvalidCredentials
	}

	// TODO: Implement JWT

	return "TOKEN", nil
}

func (s *CSOService) DeactivateCSO(username string) error {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		return err
	}
	cso.Active = false
	return s.DB.Save(&cso).Error
}

func (s *CSOService) ActivateCSO(username string) error {
	cso, err := s.GetCSOByUsername(username)
	if err != nil {
		return err
	}
	cso.Active = true
	return s.DB.Save(&cso).Error
}
