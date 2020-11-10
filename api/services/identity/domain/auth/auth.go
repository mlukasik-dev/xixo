package auth

import (
	"context"

	"go.xixo.com/api/pkg/token"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Service defines authentication service.
type Service struct {
	repo       Repository
	jwtManager *token.JWTManager
	validate   *validator.Validate
}

// New returns initialized authentication service.
func New(r Repository, m *token.JWTManager, v *validator.Validate) *Service {
	return &Service{r, m, v}
}

// LoginUser .
func (s *Service) LoginUser(ctx context.Context, accountID uuid.UUID, email, password string) (string, error) {
	matches, err := s.repo.CheckUsersPassword(ctx, accountID, email, password)
	if err != nil {
		return "", err
	}
	if !matches {
		return "", ErrInvalidPassword
	}
	info, err := s.repo.FindUserInfoByEmail(ctx, accountID, email)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(false, &accountID, info.ID, info.Roles)
	if err != nil {
		return "", err
	}
	return token, nil
}

// LoginAdmin .
func (s *Service) LoginAdmin(ctx context.Context, email, password string) (string, error) {
	matches, err := s.repo.CheckAdminsPassword(ctx, email, password)
	if err != nil {
		return "", err
	}
	if !matches {
		return "", ErrInvalidPassword
	}
	info, err := s.repo.FindAdminInfoByEmail(ctx, email)
	if err != nil {
		return "", nil
	}
	token, err := s.jwtManager.Generate(true, nil, info.ID, info.Roles)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RegisterUser .
func (s *Service) RegisterUser(ctx context.Context, accountID uuid.UUID, email, password string) (string, error) {
	info, err := s.repo.FindUserInfoByEmail(ctx, accountID, email)
	if err != nil {
		return "", err
	}
	if info.Registered {
		return "", ErrAlreadyRegistered
	}
	// Update user with password
	err = s.repo.UpdateUsersPassword(ctx, accountID, email, password)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(false, &accountID, info.ID, info.Roles)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RegisterAdmin .
func (s *Service) RegisterAdmin(ctx context.Context, email, password string) (string, error) {
	info, err := s.repo.FindAdminInfoByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if info.Registered {
		return "", ErrAlreadyRegistered
	}
	// Update admin with password
	err = s.repo.UpdateAdminsPassword(ctx, email, password)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(true, nil, info.ID, info.Roles)
	if err != nil {
		return "", err
	}
	return token, nil
}
