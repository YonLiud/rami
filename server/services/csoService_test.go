package services

import (
	"rami/models"
	"rami/utils"
	"testing"
)

var csoService *CSOService

func InitiateCSOTest() {
	testDB = utils.InitiateTestDB()
	csoService = NewCSOService(testDB)
}

func TestHashPassword(t *testing.T) {
	password := utils.GenerateRandomString(8)
	hashedPassword, err := csoService.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if hashedPassword == password {
		t.Errorf("Hashed password should not match the original password")
	}
}

func TestComparePasswords(t *testing.T) {
	password := utils.GenerateRandomString(8)
	hashedPassword, err := csoService.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !csoService.ComparePasswords(hashedPassword, password) {
		t.Errorf("Password comparison failed; the hashed and original passwords should match")
	}
}

func TestCreateCSO(t *testing.T) {
	InitiateCSOTest()

	username := utils.GenerateRandomString(10)
	password := utils.GenerateRandomString(8)
	err := csoService.CreateCSO(username, password)
	if err != nil {
		t.Fatalf("Error creating CSO: %v", err)
	}

	cso, err := csoService.GetCSOByUsername(username)
	if err != nil {
		t.Fatalf("Error retrieving CSO by username: %v", err)
	}
	if cso.Username != username {
		t.Errorf("Expected username %s, got %s", username, cso.Username)
	}

	if !csoService.ComparePasswords(cso.HashedPassword, password) {
		t.Errorf("Stored password does not match the original password")
	}
}

func TestAuthenticateCSO(t *testing.T) {
	InitiateCSOTest()

	username := utils.GenerateRandomString(8)
	password := utils.GenerateRandomString(8)

	err := csoService.CreateCSO(username, password)
	if err != nil {
		t.Fatalf("Error creating CSO: %v", err)
	}

	// Test valid credentials
	valid, err := csoService.AuthenticateCSO(username, password)
	if err != nil || !valid {
		t.Errorf("Authentication failed with correct credentials")
	}

	// Test invalid password
	invalidPassword := utils.GenerateRandomString(9)
	_, err = csoService.AuthenticateCSO(username, invalidPassword)
	if err != models.ErrInvalidCredentials {
		t.Errorf("Expected error for incorrect password, got %v", err)
	}

	// Test invalid username
	invalidUsername := utils.GenerateRandomString(9)
	_, err = csoService.AuthenticateCSO(invalidUsername, password)
	if err != models.ErrInvalidCredentials {
		t.Errorf("Expected error for incorrect username, got %v", err)
	}

	// Test deactivated CSO
	err = csoService.DeactivateCSO(username)
	if err != nil {
		t.Errorf("Error deactivating CSO: %v", err)
	}
	_, err = csoService.AuthenticateCSO(username, password)
	if err != models.ErrCSOInactive {
		t.Errorf("Expected error for deactivated CSO, got %v", err)
	}

	// Reactivate CSO and verify authentication
	err = csoService.ActivateCSO(username)
	if err != nil {
		t.Errorf("Error reactivating CSO: %v", err)
	}
	valid, err = csoService.AuthenticateCSO(username, password)
	if err != nil || !valid {
		t.Errorf("Authentication should succeed for reactivated CSO")
	}
}
