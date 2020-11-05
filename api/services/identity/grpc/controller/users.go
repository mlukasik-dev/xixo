package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/grpc/transform"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *ctr) ListUsers(ctx context.Context, req *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error) {
	users, nextPageToken, err := c.usersSvc.ListUsers(req.Parent, req.PageToken, req.PageSize)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListUsersResponse{
		Users:         transform.UsersToPb(users),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (c *ctr) GetUsersCount(ctx context.Context, req *identitypb.GetUsersCountRequest) (*identitypb.GetUsersCountResponse, error) {
	count, err := c.usersSvc.Count(req.Parent)
	if err != nil {
		return nil, err
	}
	return &identitypb.GetUsersCountResponse{
		Count: count,
	}, nil
}

func (c *ctr) GetUser(ctx context.Context, req *identitypb.GetUserRequest) (*identitypb.User, error) {
	user, err := c.usersSvc.GetUser(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return transform.UserToPb(user), nil
}

func (c *ctr) CreateUser(ctx context.Context, req *identitypb.CreateUserRequest) (*identitypb.User, error) {
	user, err := c.usersSvc.CreateUser(req.Parent, transform.PbToCreateUserInput(req.User), req.InitialUser)
	if err != nil {
		return nil, err
	}
	return transform.UserToPb(user), nil
}

func (c *ctr) UpdateUser(ctx context.Context, req *identitypb.UpdateUserRequest) (*identitypb.User, error) {
	mask, err := transform.PbToUserUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := c.usersSvc.UpdateUser(
		req.User.Name, mask, transform.PbToUpdateUserInput(req.User),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return transform.UserToPb(user), nil
}

func (c *ctr) DeleteUser(ctx context.Context, req *identitypb.DeleteUserRequest) (*empty.Empty, error) {
	err := c.usersSvc.DeleteUser(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
