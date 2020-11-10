package roles

import (
	"context"
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"

	"github.com/go-playground/validator/v10"
)

// Service role's service.
type Service struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized role's service.
func New(r Repository, v *validator.Validate) *Service {
	return &Service{r, v}
}

const maxPageSize = 1000

// ListRoles .
func (s *Service) ListRoles(ctx context.Context, pageToken string, pageSize int32, query string) (roles []Role, nextPageToken string, err error) {
	if pageSize < 1 || pageSize > maxPageSize {
		return nil, "", domain.ErrPageSizeOurOfBoundaries
	}
	var c *cursor.Cursor
	// not first request
	fmt.Printf("pageToken: %s\n", pageToken)
	if pageToken != "" {
		c, err = cursor.Decode(pageToken)
		if err != nil {
			return nil, "", domain.ErrInvalidPageToken
		}
	}
	filter, err := NewFilter(query)
	if err != nil {
		return nil, "", err
	}
	roles, err = s.repo.FindRoles(ctx, c, pageSize, filter)
	if err != nil {
		return nil, "", err
	}
	if int32(len(roles)) < pageSize {
		return roles, "", nil
	}
	// Generating next cursor
	last := roles[len(roles)-1]
	nextPageToken = cursor.Encode(&cursor.Cursor{
		Timestamp: last.CreatedAt,
		UUID:      last.ID,
	})
	return roles, nextPageToken, nil
}

// Count .
func (s *Service) Count(ctx context.Context) (int32, error) {
	count, err := s.repo.CountRoles(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetRole .
func (s *Service) GetRole(ctx context.Context, n string) (*Role, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	role, err := s.repo.FindRoleByID(ctx, name.RoleID)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// CreateRole .
func (s *Service) CreateRole(ctx context.Context, input *Role) (*Role, error) {
	role, err := s.repo.CreateRole(ctx, input)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// UpdateRole .
func (s *Service) UpdateRole(ctx context.Context, n string, mask *UpdateMask, input *Role) (*Role, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	role, err := s.repo.UpdateRole(ctx, name.RoleID, mask, input)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// DeleteRole .
func (s *Service) DeleteRole(ctx context.Context, n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteRole(ctx, name.RoleID)
	if err != nil {
		return err
	}
	return nil
}
