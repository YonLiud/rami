package models

import "errors"

var (
	ErrUsernameExists = errors.New("username already exists")

	ErrCSONotFound = errors.New("CSO not found")

	ErrInvalidCredentials = errors.New("invalid Credentials")

	ErrInactiveCSO = errors.New("CSO is inactive")
)
