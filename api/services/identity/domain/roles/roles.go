package roles

import (
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"
)

const maxPageSize = 1000

func (s *svc) ListRoles(pageToken string, pageSize int32, query string) (roles []*Role, nextPageToken string, err error) {
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
	roles, err = s.repo.FindRoles(c, pageSize, filter)
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

func (s *svc) Count() (int32, error) {
	count, err := s.repo.CountRoles()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *svc) GetRole(n string) (*Role, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	role, err := s.repo.FindRoleByID(name.RoleID)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *svc) CreateRole(input *CreateRoleInput) (*Role, error) {
	role, err := s.repo.CreateRole(input)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *svc) UpdateRole(n string, mask *UpdateMask, input *UpdateRoleInput) (*Role, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	role, err := s.repo.UpdateRole(name.RoleID, mask, input)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *svc) DeleteRole(n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteRole(name.RoleID)
	if err != nil {
		return err
	}
	return nil
}
