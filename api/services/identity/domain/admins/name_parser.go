package admins

import (
	"fmt"
	"regexp"
	"strings"

	"go.xixo.com/api/services/identity/domain"
)

// https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go
const uuidPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

var validName = regexp.MustCompile("^" + "admins" + "/" + uuidPattern + "$")

// Name .
type Name struct {
	AdminID string
}

func (n Name) String() string {
	return "admins/" + n.AdminID
}

// ParseResourceName .
func ParseResourceName(name string) (*Name, error) {
	if !validName.MatchString(name) {
		return nil, fmt.Errorf("admin %w", domain.ErrInvalidName)
	}
	s := strings.Split(name, "/")
	return &Name{
		AdminID: s[1],
	}, nil
}

// ParseCollectionName .
func ParseCollectionName() {
}
