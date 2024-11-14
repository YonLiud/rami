package models

type Visitor struct {
	Name                    string
	CredentialsNumber       string
	CredentialType          string
	VehiclePlate            string
	Association             string
	Inviter                 string
	Purpose                 string
	EntryApproval           bool
	EntryExpiry             int
	SecurityResponse        string
	ClearanceLevel          string
	ClearanceExpiry         int
	SecurityOfficerApproval bool
	Notes                   string
	Inside                  bool
}
