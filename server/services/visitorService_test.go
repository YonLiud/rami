package services

import (
	"testing"

	"rami/utils"
)

var visitorService *VisitorService

func InitiateVisitorTest() {
	testDB := utils.InitiateTestDB()
	visitorService = NewVisitorService(testDB)
}

func TestCreateVisitor(t *testing.T) {
	InitiateVisitorTest()

	demoVisitor := utils.GenerateRandomVisitor()
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

	demoVisitor := utils.GenerateRandomVisitor()
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

	demoVisitor := utils.GenerateRandomVisitor()
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
