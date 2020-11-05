package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/grpc/transform"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *ctr) ListAccounts(ctx context.Context, req *accountpb.ListAccountsRequest) (*accountpb.ListAccountsResponse, error) {
	accs, next, err := c.accountsSvc.ListAccounts(req.PageToken, req.PageSize)
	if errors.Is(err, accounts.ErrPageSizeOurOfBoundaries) || errors.Is(err, accounts.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &accountpb.ListAccountsResponse{
		Accounts:      transform.AccountsToPb(accs),
		NextPageToken: next,
	}
	return res, nil
}

func (c *ctr) GetAccountsCount(ctx context.Context, req *accountpb.GetAccountsCountRequest) (*accountpb.GetAccountsCountResponse, error) {
	count, err := c.accountsSvc.Count()
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountsCountResponse{Count: count}, nil
}

func (c *ctr) GetAccount(ctx context.Context, req *accountpb.GetAccountRequest) (*accountpb.Account, error) {
	account, err := c.accountsSvc.GetAccount(req.Name)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.AccountToPb(account), nil
}

func (c *ctr) CreateAccount(ctx context.Context, req *accountpb.CreateAccountRequest) (*accountpb.Account, error) {
	account, err := c.accountsSvc.CreateAccount(transform.PbToCreateAccountInput(req.Account))
	if err != nil {
		return nil, err
	}
	user := &identitypb.User{
		FirstName:   req.AccountAdmin.FirstName,
		LastName:    req.AccountAdmin.LastName,
		Email:       req.AccountAdmin.Email,
		PhoneNumber: req.AccountAdmin.PhoneNumber,
	}
	_, err = c.identitySvcClient.CreateUser(ctx, &identitypb.CreateUserRequest{
		Parent:      account.Name(),
		User:        user,
		InitialUser: true,
	})
	if err != nil {
		return nil, err
	}
	return transform.AccountToPb(account), nil
}

func (c *ctr) UpdateAccount(ctx context.Context, req *accountpb.UpdateAccountRequest) (*accountpb.Account, error) {
	mask, err := transform.PbToUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	account, err := c.accountsSvc.UpdateAccount(
		req.Account.Name, mask, transform.PbToUpdateAccountInput(req.Account),
	)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.AccountToPb(account), nil
}

func (c *ctr) DeleteAccount(ctx context.Context, req *accountpb.DeleteAccountRequest) (*empty.Empty, error) {
	err := c.accountsSvc.DeleteAccount(req.Name)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
