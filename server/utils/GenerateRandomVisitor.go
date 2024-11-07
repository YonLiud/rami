package utils

import (
	"rami/models"

	"github.com/brianvoe/gofakeit"
)

func GenerateRandomVisitor() models.Visitor {
	return models.Visitor{
		Name:                    gofakeit.Name(),                                           // Random Name
		CredentialsNumber:       GenerateRandomString(10),                                  // Random Credentials Number
		CredentialType:          GenerateRandomChoice([]string{"ID", "Passport", "Hoger"}), // Random Credential Type
		VehiclePlate:            GenerateRandomString(7),                                   // Random Vehicle Plate
		Association:             gofakeit.Company(),                                        // Random Company
		Inviter:                 gofakeit.Name(),                                           // Random Inviter
		Purpose:                 gofakeit.Sentence(3),                                      // Random Purpose
		EntryApproval:           true,                                                      // Fixed value
		EntryExpriry:            GenerateRandomTimestamp(),                                 // Random Timestamp
		SecurityResponse:        "Full body scan",                                          // Fixed value
		ClearanceLevel:          "High",                                                    // Fixed value
		ClearanceExpiry:         GenerateRandomTimestamp(),                                 // Random Timestamp
		SecurityOfficerApproval: true,                                                      // Fixed value
		Notes:                   gofakeit.Sentence(5),                                      // Random Notes
	}
}
