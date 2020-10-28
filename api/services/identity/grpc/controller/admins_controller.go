package controller

import (
	"context"
	"errors"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/grpc/marshaller"
	"go.xixo.com/protobuf/identitypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type adminsCtr struct {
	adminsSvc admins.Service
}

// NewAdminsController returns initialized user's gRPC controller
func NewAdminsController(adminsSvc admins.Service) identitypb.AdminsServer {
	return &adminsCtr{adminsSvc}
}

func (ctr *adminsCtr) ListAdmins(ctx context.Context, req *identitypb.ListAdminsRequest) (*identitypb.ListAdminsResponse, error) {
	admins, nextPageToken, err := ctr.adminsSvc.ListAdmins(req.PageToken, req.PageSize)
	if errors.Is(err, domain.ErrInvalidPageToken) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, err
	}
	res := &identitypb.ListAdminsResponse{
		Admins:        marshaller.AdminsToPb(admins),
		NextPageToken: nextPageToken,
	}
	return res, nil
}

func (ctr *adminsCtr) GetAdminsCount(ctx context.Context, req *identitypb.GetAdminsCountRequest) (*identitypb.GetAdminsCountResponse, error) {
	count, err := ctr.adminsSvc.Count()
	if err != nil {
		return nil, err
	}
	return &identitypb.GetAdminsCountResponse{
		Count: count,
	}, nil
}

func (ctr *adminsCtr) GetAdmin(ctx context.Context, req *identitypb.GetAdminRequest) (*identitypb.Admin, error) {
	user, err := ctr.adminsSvc.GetAdmin(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.AdminToPb(user), nil
}

func (ctr *adminsCtr) CreateAdmin(ctx context.Context, req *identitypb.CreateAdminRequest) (*identitypb.Admin, error) {
	user, err := ctr.adminsSvc.CreateAdmin(marshaller.PbToCreateAdminInput(req.Admin))
	if err != nil {
		return nil, err
	}
	return marshaller.AdminToPb(user), nil
}

func (ctr *adminsCtr) UpdateAdmin(ctx context.Context, req *identitypb.UpdateAdminRequest) (*identitypb.Admin, error) {
	mask, err := marshaller.PbToAdminUpdateMask(req.UpdateMask)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := ctr.adminsSvc.UpdateAdmin(
		req.Admin.Name, mask, marshaller.PbToUpdateAdminInput(req.Admin),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return marshaller.AdminToPb(user), nil
}

func (ctr *adminsCtr) DeleteAdmin(ctx context.Context, req *identitypb.DeleteAdminRequest) (*empty.Empty, error) {
	err := ctr.adminsSvc.DeleteAdmin(req.Name)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
