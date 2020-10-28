package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/api/services/identity/grpc/marshaller"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type usersCtr struct {
	usersSvc users.Service
}

// NewUsersController returns initialized user's gRPC controller
func NewUsersController(usersSvc users.Service) identitypb.UsersServer {
	return &usersCtr{usersSvc}
}

func (ctr *usersCtr) ListUsers(ctx context.Context, req *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error) {
	users, nextPageToken, err := ctr.usersSvc.ListUsers(req.Parent, req.PageToken, req.PageSize)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListUsersResponse{
		Users:         marshaller.UsersToPb(users),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (ctr *usersCtr) GetUsersCount(ctx context.Context, req *identitypb.GetUsersCountRequest) (*identitypb.GetUsersCountResponse, error) {
	count, err := ctr.usersSvc.Count(req.Parent)
	if err != nil {
		return nil, err
	}
	return &identitypb.GetUsersCountResponse{
		Count: count,
	}, nil
}

func (ctr *usersCtr) GetUser(ctx context.Context, req *identitypb.GetUserRequest) (*identitypb.User, error) {
	user, err := ctr.usersSvc.GetUser(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return marshaller.UserToPb(user), nil
}

func (ctr *usersCtr) CreateUser(ctx context.Context, req *identitypb.CreateUserRequest) (*identitypb.User, error) {
	user, err := ctr.usersSvc.CreateUser(req.Parent, marshaller.PbToCreateUserInput(req.User), req.InitialUser)
	if err != nil {
		return nil, err
	}
	return marshaller.UserToPb(user), nil
}

func (ctr *usersCtr) UpdateUser(ctx context.Context, req *identitypb.UpdateUserRequest) (*identitypb.User, error) {
	mask, err := marshaller.PbToUserUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := ctr.usersSvc.UpdateUser(
		req.User.Name, mask, marshaller.PbToUpdateUserInput(req.User),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return marshaller.UserToPb(user), nil
}

func (ctr *usersCtr) DeleteUser(ctx context.Context, req *identitypb.DeleteUserRequest) (*empty.Empty, error) {
	err := ctr.usersSvc.DeleteUser(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
