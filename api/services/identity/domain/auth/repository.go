package auth

// Repository .
type Repository interface {
	CheckUsersPassword(accountID, email string, plainPassword string) (bool, error)
	SetUsersPassword(accountID, email string, plainPassword string) error
	CheckAdminsPassword(email string, plainPassword string) (bool, error)
	SetAdminsPassword(email string, plainPassword string) error
	FindAdminInfoByEmail(email string) (*AdminInfo, error)
	FindUserInfoByEmail(accountID, email string) (*UserInfo, error)
}

// AdminInfo .
type AdminInfo struct {
	ID         string
	RoleIDs    []string
	Registered bool
}

// UserInfo .
type UserInfo struct {
	ID         string
	RoleIDs    []string
	Registered bool
}
