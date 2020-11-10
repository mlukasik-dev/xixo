package accounts

import (
	"context"

	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const maxPageSize = 1000

// Service account's service.
type Service struct {
	repo     Repository
	logger   *zap.Logger
	validate *validator.Validate
}

// New returns initialized account's service instance
func New(r Repository, l *zap.Logger, v *validator.Validate) *Service {
	return &Service{r, l, v}
}

// ListAccounts .
func (s *Service) ListAccounts(ctx context.Context, pageToken string, pageSize int32) (accounts []Account, nextPageToken string, err error) {
	if pageSize < 1 || pageSize > maxPageSize {
		return nil, "", ErrPageSizeOurOfBoundaries
	}
	var c *cursor.Cursor
	// not first request
	if pageToken != "" {
		c, err = cursor.Decode(pageToken)
		if err != nil {
			return nil, "", ErrInvalidPageToken
		}
	}
	accounts, err = s.repo.FindAccounts(ctx, c, pageSize)
	if err != nil {
		return nil, "", err
	}
	if int32(len(accounts)) < pageSize {
		return accounts, "", nil
	}
	// Generating next cursor
	last := accounts[len(accounts)-1]
	nextPageToken = cursor.Encode(&cursor.Cursor{
		Timestamp: last.CreatedAt,
		UUID:      last.ID,
	})

	return accounts, nextPageToken, nil
}

// CountAccounts .
func (s *Service) CountAccounts(ctx context.Context) (int32, error) {
	count, err := s.repo.CountAccounts(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAccount .
func (s *Service) GetAccount(ctx context.Context, n string) (*Account, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, ErrInvalidResourceName
	}
	account, err := s.repo.FindAccountByID(ctx, name.AccountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateAccount .
func (s *Service) CreateAccount(ctx context.Context, input *Account) (*Account, error) {
	account, err := s.repo.CreateAccount(ctx, input)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// UpdateAccount .
func (s *Service) UpdateAccount(ctx context.Context, n string, mask *UpdateMask, input *Account) (*Account, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, ErrInvalidResourceName
	}
	account, err := s.repo.UpdateAccount(ctx, name.AccountID, mask, input)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// DeleteAccount .
func (s *Service) DeleteAccount(ctx context.Context, n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return ErrInvalidResourceName
	}
	return s.repo.DeleteAccount(ctx, name.AccountID)
}
