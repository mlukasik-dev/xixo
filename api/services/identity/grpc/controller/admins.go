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

func (c *ctr) ListAdmins(ctx context.Context, req *identitypb.ListAdminsRequest) (*identitypb.ListAdminsResponse, error) {
	admins, nextPageToken, err := c.adminsSvc.ListAdmins(req.PageToken, req.PageSize)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListAdminsResponse{
		Admins:        transform.AdminsToPb(admins),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (c *ctr) GetAdminsCount(ctx context.Context, req *identitypb.GetAdminsCountRequest) (*identitypb.GetAdminsCountResponse, error) {
	count, err := c.adminsSvc.Count()
	if err != nil {
		return nil, err
	}
	return &identitypb.GetAdminsCountResponse{
		Count: count,
	}, nil
}

func (c *ctr) GetAdmin(ctx context.Context, req *identitypb.GetAdminRequest) (*identitypb.Admin, error) {
	user, err := c.adminsSvc.GetAdmin(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.AdminToPb(user), nil
}

func (c *ctr) CreateAdmin(ctx context.Context, req *identitypb.CreateAdminRequest) (*identitypb.Admin, error) {
	user, err := c.adminsSvc.CreateAdmin(transform.PbToAdmin(req.Admin))
	if err != nil {
		return nil, err
	}
	return transform.AdminToPb(user), nil
}

func (c *ctr) UpdateAdmin(ctx context.Context, req *identitypb.UpdateAdminRequest) (*identitypb.Admin, error) {
	mask, err := transform.PbToAdminUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := c.adminsSvc.UpdateAdmin(
		req.Admin.Name, mask, transform.PbToAdmin(req.Admin),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return transform.AdminToPb(user), nil
}

func (c *ctr) DeleteAdmin(ctx context.Context, req *identitypb.DeleteAdminRequest) (*empty.Empty, error) {
	err := c.adminsSvc.DeleteAdmin(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
