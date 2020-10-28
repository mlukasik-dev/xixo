package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/grpc/marshaller"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type accountsCtr struct {
	accountsSvc accounts.Service
	usersClient identitypb.UsersClient
}

// NewAccountsController creates new account's gRPC controller
func NewAccountsController(accountsSvc accounts.Service, usersClient identitypb.UsersClient) accountpb.AccountsServer {
	return &accountsCtr{accountsSvc, usersClient}
}

func (ctr *accountsCtr) ListAccounts(ctx context.Context, req *accountpb.ListAccountsRequest) (*accountpb.ListAccountsResponse, error) {
	accs, next, err := ctr.accountsSvc.ListAccounts(req.PageToken, req.PageSize)
	if errors.Is(err, accounts.ErrPageSizeOurOfBoundaries) || errors.Is(err, accounts.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &accountpb.ListAccountsResponse{
		Accounts:      marshaller.AccountsToPb(accs),
		NextPageToken: next,
	}
	return res, nil
}

func (ctr *accountsCtr) GetAccountsCount(ctx context.Context, req *accountpb.GetAccountsCountRequest) (*accountpb.GetAccountsCountResponse, error) {
	count, err := ctr.accountsSvc.Count()
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountsCountResponse{Count: count}, nil
}

func (ctr *accountsCtr) GetAccount(ctx context.Context, req *accountpb.GetAccountRequest) (*accountpb.Account, error) {
	account, err := ctr.accountsSvc.GetAccount(req.Name)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.AccountToPb(account), nil
}

func (ctr *accountsCtr) CreateAccount(ctx context.Context, req *accountpb.CreateAccountRequest) (*accountpb.Account, error) {
	account, err := ctr.accountsSvc.CreateAccount(marshaller.PbToCreateAccountInput(req.Account))
	if err != nil {
		return nil, err
	}
	user := &identitypb.User{
		FirstName:   req.AccountAdmin.FirstName,
		LastName:    req.AccountAdmin.LastName,
		Email:       req.AccountAdmin.Email,
		PhoneNumber: req.AccountAdmin.PhoneNumber,
	}
	_, err = ctr.usersClient.CreateUser(ctx, &identitypb.CreateUserRequest{
		Parent:      "accounts/" + account.ID,
		User:        user,
		InitialUser: true,
	})
	if err != nil {
		return nil, err
	}
	return marshaller.AccountToPb(account), nil
}

func (ctr *accountsCtr) UpdateAccount(ctx context.Context, req *accountpb.UpdateAccountRequest) (*accountpb.Account, error) {
	mask, err := marshaller.PbToUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	account, err := ctr.accountsSvc.UpdateAccount(
		req.Account.Name, mask, marshaller.PbToUpdateAccountInput(req.Account),
	)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.AccountToPb(account), nil
}

func (ctr *accountsCtr) DeleteAccount(ctx context.Context, req *accountpb.DeleteAccountRequest) (*empty.Empty, error) {
	err := ctr.accountsSvc.DeleteAccount(req.Name)
	if errors.Is(err, accounts.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
