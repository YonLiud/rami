package services

import (
	"rami/models"
	"testing"

	"github.com/brianvoe/gofakeit"
)

var visitorService *VisitorService

func InitiateVisitorTest() {
	testDB := InitiateTestDB()
	visitorService = NewVisitorService(testDB)
}

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
	InitiateVisitorTest()

	demoVisitor := GenerateRandomVisitor()
	err := visitorService.CreateVisitor(&demoVisitor)
	if err != nil {
		t.Fatalf("Error creating visitor: %v", err)
	}

	visitor, err := visitorService.GetVisitorByCredentialsNumber(demoVisitor.CredentialsNumber)
	if err != nil {
		t.Fatalf("Error retrieving visitor: %v", err)
	}
	if visitor.Name != demoVisitor.Name {
		t.Errorf("Expected name %s, got %s", demoVisitor.Name, visitor.Name)
	}
}

func TestGetAllVisitors(t *testing.T) {
	InitiateVisitorTest()

	demoVisitor := GenerateRandomVisitor()
	err := visitorService.CreateVisitor(&demoVisitor)
	if err != nil {
		t.Fatalf("Error creating visitor: %v", err)
	}

	visitors, err := visitorService.GetAllVisitors()
	if err != nil {
		t.Fatalf("Error retrieving visitors: %v", err)
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
	InitiateVisitorTest()

	demoVisitor := GenerateRandomVisitor()
	err := visitorService.CreateVisitor(&demoVisitor)
	if err != nil {
		t.Fatalf("Error creating visitor: %v", err)
	}

	visitor, err := visitorService.GetVisitorByCredentialsNumber(demoVisitor.CredentialsNumber)
	if err != nil {
		t.Fatalf("Error retrieving visitor: %v", err)
	}

	if visitor.CredentialsNumber != demoVisitor.CredentialsNumber {
		t.Errorf("Expected credentials number %s, got %s", demoVisitor.CredentialsNumber, visitor.CredentialsNumber)
	}
	if visitor.Name != demoVisitor.Name {
		t.Errorf("Expected name %s, got %s", demoVisitor.Name, visitor.Name)
	}
}
