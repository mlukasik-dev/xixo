package auth

import (
	"go.xixo.com/api/pkg/token"

	"github.com/go-playground/validator/v10"
)

// Service defines authentication service
type Service interface {
	LoginUser(accountID, email, password string) (string, error)
	LoginAdmin(email, password string) (string, error)
	RegisterUser(accountID, email, password string) (string, error)
	RegisterAdmin(email, password string) (string, error)
}

type svc struct {
	repo       Repository
	jwtManager *token.JWTManager
	validate   *validator.Validate
}

// New returns initialized authentication service
func New(r Repository, m *token.JWTManager, v *validator.Validate) Service {
	return &svc{r, m, v}
}
