package controller

import (
	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/protobuf/identitypb"
)

type ctr struct {
	*identitypb.UnimplementedIdentityServiceServer
	authSvc   *auth.Service
	adminsSvc *admins.Service
	usersSvc  *users.Service
	rolesSvc  *roles.Service
}

// Services .
type Services struct {
	AuthSvc   *auth.Service
	AdminsSvc *admins.Service
	UsersSvc  *users.Service
	RolesSvc  *roles.Service
}

// New returns initialized gRPC controller
func New(s *Services) identitypb.IdentityServiceServer {
	return &ctr{
		UnimplementedIdentityServiceServer: &identitypb.UnimplementedIdentityServiceServer{},
		authSvc:                            s.AuthSvc,
		adminsSvc:                          s.AdminsSvc,
		usersSvc:                           s.UsersSvc,
		rolesSvc:                           s.RolesSvc,
	}
}
