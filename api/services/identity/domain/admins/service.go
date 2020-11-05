package admins

import "github.com/go-playground/validator/v10"

// Service admin's service
type Service interface {
	ListAdmins(pageToken string, pageSize int32) ([]*Admin, string, error)
	Count() (int32, error)
	GetAdmin(name string) (*Admin, error)
	CreateAdmin(input *Admin) (*Admin, error)
	UpdateAdmin(name string, mask *UpdateMask, input *Admin) (*Admin, error)
	DeleteAdmin(name string) error
}

type svc struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized admin's service
func New(r Repository, v *validator.Validate) Service {
	return &svc{r, v}
}
