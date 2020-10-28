package admins

import "errors"

// ErrAcountIDsNotMatch .
var ErrAcountIDsNotMatch = errors.New("claimed accountID and real account does not match")
