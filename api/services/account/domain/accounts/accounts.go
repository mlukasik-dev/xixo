package accounts

import (
	"go.xixo.com/api/pkg/cursor"
)

const maxPageSize = 1000

func (s *svc) ListAccounts(pageToken string, pageSize int32) (accounts []*Account, nextPageToken string, err error) {
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
	accounts, err = s.repo.FindAccounts(c, pageSize)
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

func (s *svc) Count() (int32, error) {
	count, err := s.repo.Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *svc) GetAccount(n string) (*Account, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, ErrInvalidResourceName
	}
	account, err := s.repo.FindAccountByID(name.AccountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *svc) CreateAccount(input *CreateAccountInput) (*Account, error) {
	err := input.Validate(s.validate)
	if err != nil {
		return nil, err
	}
	account, err := s.repo.CreateAccount(input)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *svc) UpdateAccount(n string, mask *UpdateMask, input *UpdateAccountInput) (*Account, error) {
	name, err := ParseResourceName(n)
	if err != nil {
		return nil, ErrInvalidResourceName
	}
	err = input.Validate(s.validate)
	if err != nil {
		return nil, err
	}
	account, err := s.repo.UpdateAccount(name.AccountID, mask, input)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *svc) DeleteAccount(n string) error {
	name, err := ParseResourceName(n)
	if err != nil {
		return ErrInvalidResourceName
	}
	return s.repo.DeleteAccount(name.AccountID)
}
