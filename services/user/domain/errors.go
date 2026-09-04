package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyUsed   = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidCategory    = errors.New("one or more categories do not exist")
	ErrForbidden          = errors.New("forbidden")
)
