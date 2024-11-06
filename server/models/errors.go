package models

import "errors"

var ErrUsernameExists = errors.New("username already in use")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrCSOInactive = errors.New("CSO is inactive")
