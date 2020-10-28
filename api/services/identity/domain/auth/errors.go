package auth

import "errors"

// ErrInvalidPassword .
var ErrInvalidPassword = errors.New("invalid password")

// ErrNoPassword .
var ErrNoPassword = errors.New("no password found")

// ErrAlreadyRegistered .
var ErrAlreadyRegistered = errors.New("already registered")
