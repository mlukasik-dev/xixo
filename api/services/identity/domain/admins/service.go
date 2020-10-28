package admins

import "github.com/go-playground/validator/v10"

// Service admin's service
type Service interface {
	ListAdmins(pageToken string, pageSize int32) ([]*Admin, string, error)
	Count() (int32, error)
	GetAdmin(id string) (*Admin, error)
	CreateAdmin(input *CreateAdminInput) (*Admin, error)
	UpdateAdmin(id string, mask *UpdateMask, input *UpdateAdminInput) (*Admin, error)
	DeleteAdmin(id string) error
}

type svc struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized admin's service
func New(r Repository, v *validator.Validate) Service {
	return &svc{r, v}
}
