package admins

import (
	"context"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"

	"github.com/go-playground/validator/v10"
)

const maxPageSize = 1000

// Service admin's service.
type Service struct {
	repo     Repository
	validate *validator.Validate
}

// New returns initialized admin's service.
func New(r Repository, v *validator.Validate) *Service {
	return &Service{r, v}
}

// ListAdmins .
func (s *Service) ListAdmins(ctx context.Context, pageToken string, pageSize int32) (admins []Admin, nextPageToken string, err error) {
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
	admins, err = s.repo.FindAdmins(ctx, c, pageSize)
	if err != nil {
		return nil, "", err
	}
	if int32(len(admins)) < pageSize {
		return admins, "", nil
	}
	// Generating next cursor
	last := admins[len(admins)-1]
	nextPageToken = cursor.Encode(&cursor.Cursor{
		Timestamp: last.CreatedAt,
		UUID:      last.ID,
	})
	return admins, nextPageToken, nil
}

// Count .
func (s *Service) Count(ctx context.Context) (int32, error) {
	count, err := s.repo.CountAdmins(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAdmin .
func (s *Service) GetAdmin(ctx context.Context, n string) (*Admin, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	admin, err := s.repo.FindAdminByID(ctx, name.AdminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// CreateAdmin .
func (s *Service) CreateAdmin(ctx context.Context, adminInput *Admin) (admin *Admin, err error) {
	admin, err = s.repo.CreateAdmin(ctx, adminInput)
	if err != nil {
		return nil, err
	}
	return
}

// UpdateAdmin .
func (s *Service) UpdateAdmin(ctx context.Context, n string, mask *UpdateMask, adminInput *Admin) (*Admin, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	admin, err := s.repo.UpdateAdmin(ctx, name.AdminID, mask, adminInput)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// DeleteAdmin .
func (s *Service) DeleteAdmin(ctx context.Context, n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteAdmin(ctx, name.AdminID)
	if err != nil {
		return err
	}
	return nil
}
