package users

import (
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"
)

const maxPageSize = 1000

func (s *svc) ListUsers(parent string, pageToken string, pageSize int32) (users []*User, nextPageToken string, err error) {
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
	users, err = s.repo.FindUsers(name.AccountID, c, pageSize)
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

func (s *svc) Count(parent string) (int32, error) {
	name, err := ParseCollectionName(parent)
	if err != nil {
		return 0, err
	}
	count, err := s.repo.CountUsers(name.AccountID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *svc) GetUser(n string) (*User, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindUserByID(name.AccountID, name.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *svc) CreateUser(parent string, input *CreateUserInput, initial bool) (user *User, err error) {
	err = input.Validate(s.validate)
	if err != nil {
		return nil, err
	}
	name, err := ParseCollectionName(parent)
	if err != nil {
		return nil, err
	}
	if initial {
		user, err = s.repo.CreateAccountAdmin(name.AccountID, input)
	} else {
		user, err = s.repo.CreateUser(name.AccountID, input)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *svc) UpdateUser(n string, mask *UpdateMask, input *UpdateUserInput) (*User, error) {
	err := input.Validate(s.validate)
	if err != nil {
		return nil, err
	}
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, err
	}
	fmt.Printf("accountID: %s, userID: %s\n", name.AccountID, name.UserID)
	user, err := s.repo.UpdateUser(name.AccountID, name.UserID, mask, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *svc) DeleteUser(n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return err
	}
	err = s.repo.DeleteUser(name.AccountID, name.UserID)
	if err != nil {
		return err
	}
	return nil
}
