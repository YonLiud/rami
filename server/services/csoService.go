package services

import (
	"log"
	"rami/database"
	"rami/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func DeactivateCSO(username string) error {
	db := database.GetDB()

	cso, err := GetCSOByUsername(username)
	if err != nil {
		return err
	}

	cso.Active = false
	return db.Save(&cso).Error
}

func ActivateCSO(username string) error {
	db := database.GetDB()

	cso, err := GetCSOByUsername(username)
	if err != nil {
		return err
	}

	cso.Active = true
	return db.Save(&cso).Error
}

func CreateCSO(username string, password string) error {
	db := database.GetDB()

	var existingCSO models.CSO
	err := db.Where("username = ?", username).First(&existingCSO).Error
	if err != gorm.ErrRecordNotFound {
		return models.ErrUsernameExists
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	cso := models.CSO{
		Username:       username,
		HashedPassword: hashedPassword,
		Active:         true,
	}

	return db.Create(&cso).Error
}

func GetCSOByUsername(username string) (models.CSO, error) {
	db := database.GetDB()

	var cso models.CSO
	err := db.Where("username = ?", username).First(&cso).Error
	if err != nil {
		return cso, err
	}
	return cso, nil
}

func GetAllActiveCSOs() ([]models.CSO, error) {
	db := database.GetDB()

	var csos []models.CSO
	err := db.Where("active = ?", true).Find(&csos).Error
	return csos, err
}

func AuthenticateCSO(username, password string) (bool, error) {
	// check if the user exists
	cso, err := GetCSOByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, models.ErrInvalidCredentials
		}
		return false, err
	}

	if !cso.Active {
		return false, models.ErrCSOInactive
	}

	if ComparePasswords(cso.HashedPassword, password) {
		return true, nil
	}

	return false, models.ErrInvalidCredentials
}
