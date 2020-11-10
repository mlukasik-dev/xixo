package users

import (
	"context"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"

	"github.com/go-playground/validator/v10"
)

const maxPageSize = 1000

// Service user's service.
type Service struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized user's service.
func New(r Repository, v *validator.Validate) *Service {
	return &Service{r, v}
}

// ListUsers .
func (s *Service) ListUsers(ctx context.Context, parent string, pageToken string, pageSize int32) (users []User, nextPageToken string, err error) {
	if pageSize < 1 || pageSize > maxPageSize {
		return nil, "", domain.ErrPageSizeOurOfBoundaries
	}
	var c *cursor.Cursor
	// not first request
	if pageToken != "" {
		c, err = cursor.Decode(pageToken)
		if err != nil {
			return nil, "", domain.ErrInvalidPageToken
		}
	}
	name, err := ParseCollectionName(parent)
	if err != nil {
		return nil, "", err
	}
	users, err = s.repo.FindUsers(ctx, name.AccountID, c, pageSize)
	if err != nil {
		return nil, "", err
	}
	if int32(len(users)) < pageSize {
		return users, "", nil
	}
	// Generating next cursor
	last := users[len(users)-1]
	nextPageToken = cursor.Encode(&cursor.Cursor{
		Timestamp: last.CreatedAt,
		UUID:      last.ID,
	})
	return users, nextPageToken, nil
}

// Count .
func (s *Service) Count(ctx context.Context, parent string) (int32, error) {
	name, err := ParseCollectionName(parent)
	if err != nil {
		return 0, err
	}
	count, err := s.repo.CountUsers(ctx, name.AccountID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetUser .
func (s *Service) GetUser(ctx context.Context, n string) (*User, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindUserByID(ctx, name.AccountID, name.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser .
func (s *Service) CreateUser(ctx context.Context, parent string, input *User, initial bool) (user *User, err error) {
	name, err := ParseCollectionName(parent)
	if err != nil {
		return nil, err
	}
	if initial {
		user, err = s.repo.CreateAccountAdmin(ctx, name.AccountID, input)
	} else {
		user, err = s.repo.CreateUser(ctx, name.AccountID, input)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser .
func (s *Service) UpdateUser(ctx context.Context, n string, mask *UpdateMask, input *User) (*User, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.UpdateUser(ctx, name.AccountID, name.UserID, mask, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser .
func (s *Service) DeleteUser(ctx context.Context, n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteUser(ctx, name.AccountID, name.UserID)
	if err != nil {
		return err
	}
	return nil
}
