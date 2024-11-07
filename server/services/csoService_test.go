package services

import (
	"rami/models"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := generateRandomString(8)
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if hashedPassword == password {
		t.Errorf("Hashed password is the same as the password")
	}
}

func TestComparePasswords(t *testing.T) {
	password := generateRandomString(8)
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !ComparePasswords(hashedPassword, password) {
		t.Errorf("Passwords should match")
	}
}

func TestCreateCSO(t *testing.T) {
	InitiateTestDB()

	username := generateRandomString(10)
	password := generateRandomString(8)
	err := CreateCSO(username, password)
	if err != nil {
		t.Errorf("Error creating CSO: %v", err)
		t.FailNow()
	}

	cso, err := GetCSOByUsername(username)
	if err != nil {
		t.Errorf("Error getting CSO by username: %v", err)
	}
	if cso.Username != username {
		t.Errorf("Username does not match")
	}

	if !ComparePasswords(cso.HashedPassword, password) {
		t.Errorf("Passwords should match")
	}
}

func TestAuthenticateCSO(t *testing.T) {
	InitiateTestDB()

	username := generateRandomString(8)
	password := generateRandomString(8)

	err := CreateCSO(username, password)
	if err != nil {
		t.Errorf("Error creating CSO: %v", err)
	}

	valid, err := AuthenticateCSO(username, password)
	if err != nil {
		t.Errorf("Error authenticating CSO: %v", err)
	}
	if !valid {
		t.Errorf("Authentication should succeed with correct password")
	}

	invalidPassword := generateRandomString(9)
	_, err = AuthenticateCSO(username, invalidPassword)
	if err != nil {
		if err != models.ErrInvalidCredentials {
			t.Errorf("Error authenticating CSO: %v", err)
		}
	} else {
		t.Errorf("Authentication should fail with incorrect password")
	}

	invalidUsername := generateRandomString(9)
	_, err = AuthenticateCSO(invalidUsername, password)
	if err != nil {
		if err != models.ErrInvalidCredentials {
			t.Errorf("Error authenticating CSO: %v", err)
			t.FailNow()
		}
	} else {
		t.Errorf("Authentication should fail with incorrect username")

	}

	DeactivateCSO(username)
	_, err = AuthenticateCSO(username, password)
	if err != nil {
		if err != models.ErrCSOInactive {
			t.Errorf("Error authenticating CSO: %v", err)
		}
	} else {
		t.Errorf("Authentication should fail with deactivated CSO")
	}

	ActivateCSO(username)
	valid, err = AuthenticateCSO(username, password)
	if err != nil {
		t.Errorf("Error authenticating CSO: %v", err)
	}
	if !valid {
		t.Errorf("Authentication should succeed with activated CSO")
	}
}
