package models

import (
	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model
	Name                    string `json:"name"`
	CredentialsNumber       string `json:"credentialsNumber"`       // credentials number
	CredentialType          string `json:"credentialType"`          // passport, id, hoger
	VehiclePlate            string `json:"vehiclePlate"`            // vehicle plate if any
	Association             string `json:"association"`             // company / organization
	Inviter                 string `json:"inviter"`                 // person who invited the Visitor
	Purpose                 string `json:"purpose"`                 // purpose of visit
	EntryApproval           bool   `json:"entryApproval"`           // entry approval
	EntryExpriry            int    `json:"entryExpriry"`            // entry expiry UNIX
	SecurityResponse        string `json:"securityResponse"`        // what type of security measures needed upon entry
	ClearanceLevel          string `json:"clearanceLevel"`          // security clearance level
	ClearanceExpiry         int    `json:"clearanceExpiry"`         // security clearance expiry UNIX
	SecurityOfficerApproval bool   `json:"securityOfficerApproval"` // security officer approval in MOD
	Notes                   string `json:"notes"`                   // additional notes
}
