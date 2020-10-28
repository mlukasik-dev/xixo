package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authCtr struct {
	authSvc auth.Service
}

// NewAuthController returns initialized authentication gRPC controller
func NewAuthController(authSvc auth.Service) identitypb.AuthServer {
	return &authCtr{authSvc}
}

func (ctr *authCtr) Login(ctx context.Context, req *identitypb.LoginRequest) (*identitypb.Token, error) {
	var token string
	var err error
	if req.AccountId != "" {
		// Authenticate as user
		token, err = ctr.authSvc.LoginUser(req.AccountId, req.Email, req.Password)
	} else {
		// Authenticate as admin
		token, err = ctr.authSvc.LoginAdmin(req.Email, req.Password)
	}
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.PermissionDenied, "no user with such email")
	}
	if errors.Is(err, auth.ErrInvalidPassword) {
		return nil, status.Errorf(codes.PermissionDenied, "invalid password")
	}
	if errors.Is(err, auth.ErrNoPassword) {
		return nil, status.Errorf(codes.PermissionDenied, "user needs to be registered first")
	}
	if err != nil {
		return nil, err
	}
	return &identitypb.Token{AccessToken: token}, nil
}

func (ctr *authCtr) Register(ctx context.Context, req *identitypb.RegisterRequest) (*identitypb.Token, error) {
	var token string
	var err error
	if req.AccountId != "" {
		// Register as user
		token, err = ctr.authSvc.RegisterUser(req.AccountId, req.Email, req.Password)
	} else {
		// Register as admin
		token, err = ctr.authSvc.RegisterAdmin(req.Email, req.Password)
	}
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.PermissionDenied, "user with such email was not found")
	}
	if err != nil {
		return nil, err
	}
	return &identitypb.Token{AccessToken: token}, nil
}
