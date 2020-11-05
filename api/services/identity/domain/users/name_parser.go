package users

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"go.xixo.com/api/services/identity/domain"
)

// https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go
const uuidPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var validResourceName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "/users/" + uuidPattern + "$")
var validCollectionName = regexp.MustCompile("^" + "accounts/" + uuidPattern + "$")

// Name .
type Name struct {
	AccountID uuid.UUID
	UserID    uuid.UUID
}

func (n Name) String() string {
	return fmt.Sprintf("accounts/%s/users/%s", n.AccountID, n.UserID)
}

// ParseResourceName .
func ParseResourceName(name string) (*Name, error) {
	if !validResourceName.MatchString(name) {
		return nil, fmt.Errorf("user %w", domain.ErrInvalidName)
	}
	s := strings.Split(name, "/")
	accountID, err := uuid.Parse(s[1])
	userID, err := uuid.Parse(s[3])
	return &Name{
		AccountID: accountID,
		UserID:    userID,
	}, err
}

// ParseCollectionName .
func ParseCollectionName(name string) (*Name, error) {
	if !validCollectionName.MatchString(name) {
		return nil, fmt.Errorf("user %w", domain.ErrInvalidName)
	}
	s := strings.Split(name, "/")
	accountID, err := uuid.Parse(s[1])
	return &Name{
		AccountID: accountID,
	}, err
}
