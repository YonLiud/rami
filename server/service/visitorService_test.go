package service

import (
	"rami/models"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func GenerateRandomVisitor() models.Visitor {
	return models.Visitor{
		Name:                    gofakeit.Name(),                                           // Random Name
		CredentialsNumber:       generateRandomString(10),                                  // Random Credentials Number
		CredentialType:          generateRandomChoice([]string{"ID", "Passport", "Hoger"}), // Random Credential Type
		VehiclePlate:            generateRandomString(7),                                   // Random Vehicle Plate
		Association:             gofakeit.Company(),                                        // Random Company
		Inviter:                 gofakeit.Name(),                                           // Random Inviter
		Purpose:                 gofakeit.Sentence(3),                                      // Random Purpose
		EntryApproval:           true,                                                      // Fixed value
		EntryExpriry:            generateRandomTimestamp(),                                 // Random Timestamp
		SecurityResponse:        "Full body scan",                                          // Fixed value
		ClearanceLevel:          "High",                                                    // Fixed value
		ClearanceExpiry:         generateRandomTimestamp(),                                 // Random Timestamp
		SecurityOfficerApproval: true,                                                      // Fixed value
		Notes:                   gofakeit.Sentence(5),                                      // Random Notes
	}
}

func TestCreateVisitor(t *testing.T) {
	InitiateTestDB()
	demoVisitor := GenerateRandomVisitor()

	err := CreateVisitor(&demoVisitor)
	if err != nil {
		t.Errorf("Error creating visitor: %v", err)
	}

	visitor, err := GetVisitorByCredentialsNumber(demoVisitor.CredentialsNumber)
	if err != nil {
		t.Errorf("Error retrieving visitor: %v", err)
	}
	if visitor.Name != demoVisitor.Name {
		t.Errorf("Retrieved visitor name: got %v, want %v", visitor.Name, demoVisitor.Name)
	}
}

func TestGetAllVisitors(t *testing.T) {
	InitiateTestDB()
	demoVisitor := GenerateRandomVisitor()

	err := CreateVisitor(&demoVisitor)
	if err != nil {
		t.Errorf("Error creating visitor: %v", err)
	}

	visitors, err := GetAllVisitors()
	if err != nil {
		t.Errorf("Error retrieving visitors: %v", err)
	}

	if len(visitors) < 1 {
		t.Errorf("Expected at least 1 visitor, got %v", len(visitors))
	}

	found := false
	for _, visitor := range visitors {
		if visitor.CredentialsNumber == demoVisitor.CredentialsNumber {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Visitor with credentials number %v not found in all visitors", demoVisitor.CredentialsNumber)
	}
}

func TestGetVisitorByCredentialsNumber(t *testing.T) {
	InitiateTestDB()
	demoVisitor := GenerateRandomVisitor()

	err := CreateVisitor(&demoVisitor)

	if err != nil {
		t.Errorf("Error creating visitor: %v", err)
	}

	visitor, err := GetVisitorByCredentialsNumber(demoVisitor.CredentialsNumber)
	if err != nil {
		t.Errorf("Error retrieving visitor: %v", err)
	}

	if visitor.CredentialsNumber != demoVisitor.CredentialsNumber {
		t.Errorf("Retrieved visitor's credentials number: got %v, want %v", visitor.CredentialsNumber, demoVisitor.CredentialsNumber)
	}

	if visitor.Name != demoVisitor.Name {
		t.Errorf("Retrieved visitor name: got %v, want %v", visitor.Name, demoVisitor.Name)
	}
}
