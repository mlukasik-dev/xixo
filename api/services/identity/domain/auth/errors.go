package auth

import "errors"

var (
	// ErrInvalidPassword .
	ErrInvalidPassword = errors.New("invalid password")

	// ErrNoPassword .
	ErrNoPassword = errors.New("no password found")

	// ErrAlreadyRegistered .
	ErrAlreadyRegistered = errors.New("already registered")
)
