package service

import (
	"rami/models"
	"testing"
)

var demoVisitor = models.Visitor{
	Name:                    "John Doe",
	CredentialsNumber:       "ABC123456",
	CredentialType:          "passport",
	VehiclePlate:            "XYZ 1234",
	Association:             "Acme Corporation",
	Inviter:                 "Jane Smith",
	Purpose:                 "Meeting",
	EntryApproval:           true,
	EntryExpriry:            1691908800,
	SecurityResponse:        "Full body scan",
	ClearanceLevel:          "High",
	ClearanceExpiry:         1692513600,
	SecurityOfficerApproval: true,
	Notes:                   "Special instructions: escort required",
}

func TestCreateVisitor(t *testing.T) {
	InitiateTestDB()

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
