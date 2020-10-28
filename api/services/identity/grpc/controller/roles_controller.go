package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/grpc/marshaller"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rolesCtr struct {
	rolesSvc roles.Service
}

// NewRolesController returns initialized roles gRPC controller
func NewRolesController(rolesSvc roles.Service) identitypb.RolesServer {
	return &rolesCtr{rolesSvc}
}

func (ctr *rolesCtr) ListRoles(ctx context.Context, req *identitypb.ListRolesRequest) (*identitypb.ListRolesResponse, error) {
	roles, nextPageToken, err := ctr.rolesSvc.ListRoles(req.PageToken, req.PageSize, req.Filter)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListRolesResponse{
		Roles:         marshaller.RolesToPb(roles),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (ctr *rolesCtr) GetRolesCount(ctx context.Context, req *identitypb.GetRolesCountRequest) (*identitypb.GetRolesCountResponse, error) {
	count, err := ctr.rolesSvc.Count()
	if err != nil {
		return nil, err
	}
	return &identitypb.GetRolesCountResponse{
		Count: count,
	}, nil
}

func (ctr *rolesCtr) GetRole(ctx context.Context, req *identitypb.GetRoleRequest) (*identitypb.Role, error) {
	role, err := ctr.rolesSvc.GetRole(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.RoleToPb(role), nil
}

func (ctr *rolesCtr) CreateRole(ctx context.Context, req *identitypb.CreateRoleRequest) (*identitypb.Role, error) {
	role, err := ctr.rolesSvc.CreateRole(marshaller.PbToCreateRoleInput(req.Role))
	if err != nil {
		return nil, err
	}
	return marshaller.RoleToPb(role), nil
}

func (ctr *rolesCtr) UpdateRole(ctx context.Context, req *identitypb.UpdateRoleRequest) (*identitypb.Role, error) {
	mask, err := marshaller.PbToRoleUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role, err := ctr.rolesSvc.UpdateRole(
		req.Role.Name, mask, marshaller.PbToUpdateRoleInput(req.Role),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.RoleToPb(role), nil
}

func (ctr *rolesCtr) DeleteRole(ctx context.Context, req *identitypb.DeleteRoleRequest) (*empty.Empty, error) {
	err := ctr.rolesSvc.DeleteRole(req.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
