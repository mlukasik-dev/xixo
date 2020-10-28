package users

import "github.com/go-playground/validator/v10"

// Service user's service
type Service interface {
	ListUsers(parent string, pageToken string, pageSize int32) ([]*User, string, error)
	Count(parent string) (int32, error)
	GetUser(name string) (*User, error)
	CreateUser(parent string, input *CreateUserInput, initial bool) (*User, error)
	UpdateUser(name string, mask *UpdateMask, input *UpdateUserInput) (*User, error)
	DeleteUser(name string) error
}

type svc struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized user's service
func New(r Repository, v *validator.Validate) Service {
	return &svc{r, v}
}
