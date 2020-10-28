package accounts

import (
	"regexp"
	"strings"
)

// https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go
const uuidPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var validName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "$")

// Name .
type Name struct {
	AccountID string
}

func (n Name) String() string {
	return "accounts/" + n.AccountID
}

// ParseResourceName .
func ParseResourceName(name string) (*Name, error) {
	// allow empty (case of resource creation)
	if name == "" {
		return &Name{}, nil
	}
	if !validName.MatchString(name) {
		return nil, ErrInvalidResourceName
	}
	s := strings.Split(name, "/")
	return &Name{
		AccountID: s[1],
	}, nil
}

// ParseCollectionName .
func ParseCollectionName() {

}
