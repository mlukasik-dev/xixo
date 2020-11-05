package admins

import (
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"
)

const maxPageSize = 1000

func (s *svc) ListAdmins(pageToken string, pageSize int32) (admins []*Admin, nextPageToken string, err error) {
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
	admins, err = s.repo.FindAdmins(c, pageSize)
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

func (s *svc) Count() (int32, error) {
	count, err := s.repo.CountAdmins()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *svc) GetAdmin(n string) (*Admin, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	admin, err := s.repo.FindAdminByID(name.AdminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *svc) CreateAdmin(adminInput *Admin) (admin *Admin, err error) {
	admin, err = s.repo.CreateAdmin(adminInput)
	if err != nil {
		return nil, err
	}
	return
}

func (s *svc) UpdateAdmin(n string, mask *UpdateMask, adminInput *Admin) (*Admin, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	admin, err := s.repo.UpdateAdmin(name.AdminID, mask, adminInput)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *svc) DeleteAdmin(n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteAdmin(name.AdminID)
	if err != nil {
		return err
	}
	return nil
}
