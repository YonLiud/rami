package models

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Name             string `json:"name"`
	ID_Number        string `json:"id"`
	VehiclePlate     string `json:"vehicle_plate"`
	Affiliation      string `json:"affiliation"`
	Inviter          string `json:"inviter"`
	Purpose          string `json:"purpose"`
	SecurityResp     string `json:"security"`
	EntryApproved    string `json:"entry"`
	SecurityClear    string `json:"securityClear"`
	SecurityClearExp string `json:"securityClearExp"`
	SecurityMgrAppr  string `json:"securityMgrAppr"`
	ApprovalExp      string `json:"approvalExp"`
	Notes            string `json:"notes"`
}
