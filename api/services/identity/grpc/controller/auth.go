package controller

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *ctr) Login(ctx context.Context, req *identitypb.LoginRequest) (*identitypb.Token, error) {
	var token string
	var err error
	if req.AccountId != "" {
		// Authenticate as user
		accountID, err := uuid.Parse(req.AccountId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid uuid")
		}
		token, err = c.authSvc.LoginUser(accountID, req.Email, req.Password)
	} else {
		// Authenticate as admin
		token, err = c.authSvc.LoginAdmin(req.Email, req.Password)
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

func (c *ctr) Register(ctx context.Context, req *identitypb.RegisterRequest) (*identitypb.Token, error) {
	var token string
	var err error
	if req.AccountId != "" {
		// Register as user
		accountID, err := uuid.Parse(req.AccountId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid uuid")
		}
		token, err = c.authSvc.RegisterUser(accountID, req.Email, req.Password)
	} else {
		// Register as admin
		token, err = c.authSvc.RegisterAdmin(req.Email, req.Password)
	}
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.PermissionDenied, "user with such email was not found")
	}
	if err != nil {
		return nil, err
	}
	return &identitypb.Token{AccessToken: token}, nil
}
