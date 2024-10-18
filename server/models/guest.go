package models

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Name             string `json:"name"`
	ID_Number        string `json:"id" gorm:"unique"`
	VehiclePlate     string `json:"vehiclePlate"`
	Affiliation      string `json:"affiliation"`
	Inviter          string `json:"inviter"`
	Purpose          string `json:"purpose"`
	SecurityResp     string `json:"security"`
	EntryApproved    bool   `json:"entry"`
	SecurityClear    string `json:"securityClear"`
	SecurityClearExp string `json:"securityClearExp"`
	SecurityMgrAppr  bool   `json:"securityMgrAppr"`
	ApprovalExp      string `json:"approvalExp"`
	Notes            string `json:"notes"`
	IsInside         bool   `json:"isInside"`
}
