package roles

import "github.com/go-playground/validator/v10"

// Service role's service
type Service interface {
	ListRoles(pageToken string, pageSize int32, query string) ([]*Role, string, error)
	Count() (int32, error)
	GetRole(name string) (*Role, error)
	CreateRole(input *CreateRoleInput) (*Role, error)
	UpdateRole(name string, mask *UpdateMask, input *UpdateRoleInput) (*Role, error)
	DeleteRole(name string) error
}

type svc struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized role's service
func New(r Repository, v *validator.Validate) Service {
	return &svc{r, v}
}
