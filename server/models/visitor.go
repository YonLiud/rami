package models

import (
	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model
	Name                    string `json:"name" gorm:"not null"`                                        // name of the visitor
	CredentialsNumber       string `json:"credentialsNumber" gorm:"not null;uniqueIndex;default:null" ` // credentials number
	CredentialType          string `json:"credentialType" gorm:"not null"`                              // passport, id, hoger
	VehiclePlate            string `json:"vehiclePlate" gorm:"not null"`                                // vehicle plate if any
	Association             string `json:"association" gorm:"not null"`                                 // company / organization
	Inviter                 string `json:"inviter" gorm:"not null"`                                     // person who invited the Visitor
	Purpose                 string `json:"purpose" gorm:"not null"`                                     // purpose of visit
	EntryApproval           bool   `json:"entryApproval" gorm:"not null"`                               // entry approval
	EntryExpriry            int    `json:"entryExpriry" gorm:"not null"`                                // entry expiry UNIX
	SecurityResponse        string `json:"securityResponse" gorm:"not null"`                            // what type of security measures needed upon entry
	ClearanceLevel          string `json:"clearanceLevel" gorm:"not null"`                              // security clearance level
	ClearanceExpiry         int    `json:"clearanceExpiry" gorm:"not null"`                             // security clearance expiry UNIX
	SecurityOfficerApproval bool   `json:"securityOfficerApproval" gorm:"not null"`                     // security officer approval in MOD
	Notes                   string `json:"notes" gorm:"not null"`
	Inside                  bool   `json:"inside" gorm:"not null; default:false"`
}
