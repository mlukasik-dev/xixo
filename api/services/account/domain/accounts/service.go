package accounts

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Service account's service
type Service interface {
	ListAccounts(pageToken string, pageSize int32) ([]*Account, string, error)
	Count() (int32, error)
	GetAccount(name string) (*Account, error)
	CreateAccount(input *CreateAccountInput) (*Account, error)
	UpdateAccount(name string, mask *UpdateMask, input *UpdateAccountInput) (*Account, error)
	DeleteAccount(name string) error
}

type svc struct {
	repo     Repository
	logger   *zap.Logger
	validate *validator.Validate
}

// New returns initialized account's service instance
func New(r Repository, l *zap.Logger, v *validator.Validate) Service {
	return &svc{r, l, v}
}
