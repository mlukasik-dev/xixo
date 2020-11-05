package accounts

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go
const uuidPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var validName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "$")

// Name .
type Name struct {
	AccountID uuid.UUID
}

func (n Name) String() string {
	return "accounts/" + n.AccountID.String()
}

// ParseResourceName .
func ParseResourceName(name string) (*Name, error) {
	if !validName.MatchString(name) {
		return nil, ErrInvalidResourceName
	}
	s := strings.Split(name, "/")
	accountID, err := uuid.Parse(s[1])
	if err != nil {
		return nil, err
	}
	return &Name{
		AccountID: accountID,
	}, nil
}

// ParseCollectionName .
func ParseCollectionName() {

}
