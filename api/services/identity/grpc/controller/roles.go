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

func (c *ctr) ListRoles(ctx context.Context, req *identitypb.ListRolesRequest) (*identitypb.ListRolesResponse, error) {
	roles, nextPageToken, err := c.rolesSvc.ListRoles(ctx, req.PageToken, req.PageSize, req.Filter)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListRolesResponse{
		Roles:         transform.RolesToPb(roles),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (c *ctr) GetRolesCount(ctx context.Context, req *identitypb.GetRolesCountRequest) (*identitypb.GetRolesCountResponse, error) {
	count, err := c.rolesSvc.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &identitypb.GetRolesCountResponse{
		Count: count,
	}, nil
}

func (c *ctr) GetRole(ctx context.Context, req *identitypb.GetRoleRequest) (*identitypb.Role, error) {
	role, err := c.rolesSvc.GetRole(ctx, req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.RoleToPb(role), nil
}

func (c *ctr) CreateRole(ctx context.Context, req *identitypb.CreateRoleRequest) (*identitypb.Role, error) {
	role, err := c.rolesSvc.CreateRole(ctx, transform.PbToRole(req.Role))
	if err != nil {
		return nil, err
	}
	return transform.RoleToPb(role), nil
}

func (c *ctr) UpdateRole(ctx context.Context, req *identitypb.UpdateRoleRequest) (*identitypb.Role, error) {
	mask, err := transform.PbToRoleUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role, err := c.rolesSvc.UpdateRole(
		ctx, req.Role.Name, mask, transform.PbToRole(req.Role),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.RoleToPb(role), nil
}

func (c *ctr) DeleteRole(ctx context.Context, req *identitypb.DeleteRoleRequest) (*empty.Empty, error) {
	err := c.rolesSvc.DeleteRole(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
