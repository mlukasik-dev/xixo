package users

import "errors"

var (
	// ErrAcountIDsNotMatch .
	ErrAcountIDsNotMatch = errors.New("claimed accountID and real account does not match")
)
