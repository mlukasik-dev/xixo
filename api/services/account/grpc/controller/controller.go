package controller

import (
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"
)

type ctr struct {
	*accountpb.UnimplementedAccountServiceServer
	accountsSvc       *accounts.Service
	identitySvcClient identitypb.IdentityServiceClient
}

// New creates new account's gRPC controller
func New(accountsSvc *accounts.Service, usersClient identitypb.IdentityServiceClient) accountpb.AccountServiceServer {
	return &ctr{
		&accountpb.UnimplementedAccountServiceServer{},
		accountsSvc, usersClient,
	}
}
