package users

import (
	"fmt"
	"regexp"
	"strings"

	"go.xixo.com/api/services/identity/domain"
)

// https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go
const uuidPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var validResourceName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "/users/" + uuidPattern + "$")
var validCollectionName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "$")

// Name .
type Name struct {
	AccountID string
	UserID    string
}

func (n Name) String() string {
	return "accounts/" + n.AccountID + "/users/" + n.UserID
}

// ParseResourceName .
func ParseResourceName(name string) (*Name, error) {
	if !validResourceName.MatchString(name) {
		return nil, fmt.Errorf("user %w", domain.ErrInvalidName)
	}
	s := strings.Split(name, "/")
	return &Name{
		AccountID: s[1],
		UserID:    s[3],
	}, nil
}

// ParseCollectionName .
func ParseCollectionName(name string) (*Name, error) {
	if !validCollectionName.MatchString(name) {
		return nil, fmt.Errorf("user %w", domain.ErrInvalidName)
	}
	s := strings.Split(name, "/")
	return &Name{
		AccountID: s[1],
	}, nil
}
