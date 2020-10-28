package auth

func (s *svc) LoginUser(accountID, email, password string) (string, error) {
	matches, err := s.repo.CheckUsersPassword(accountID, email, password)
	if err != nil {
		return "", err
	}
	if !matches {
		return "", ErrInvalidPassword
	}
	info, err := s.repo.FindUserInfoByEmail(accountID, email)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(false, accountID, info.ID, info.RoleIDs)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *svc) LoginAdmin(email, password string) (string, error) {
	matches, err := s.repo.CheckAdminsPassword(email, password)
	if err != nil {
		return "", err
	}
	if !matches {
		return "", ErrInvalidPassword
	}
	info, err := s.repo.FindAdminInfoByEmail(email)
	if err != nil {
		return "", nil
	}
	token, err := s.jwtManager.Generate(true, "", info.ID, info.RoleIDs)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *svc) RegisterUser(accountID, email, password string) (string, error) {
	info, err := s.repo.FindUserInfoByEmail(accountID, email)
	if err != nil {
		return "", err
	}
	if info.Registered {
		return "", ErrAlreadyRegistered
	}
	// Update user with password
	err = s.repo.SetUsersPassword(accountID, email, password)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(false, accountID, info.ID, info.RoleIDs)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *svc) RegisterAdmin(email, password string) (string, error) {
	info, err := s.repo.FindAdminInfoByEmail(email)
	if err != nil {
		return "", err
	}
	if info.Registered {
		return "", ErrAlreadyRegistered
	}
	// Update admin with password
	err = s.repo.SetAdminsPassword(email, password)
	if err != nil {
		return "", err
	}
	token, err := s.jwtManager.Generate(true, "", info.ID, info.RoleIDs)
	if err != nil {
		return "", err
	}
	return token, nil
}
