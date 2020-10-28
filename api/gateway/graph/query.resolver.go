package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vektah/gqlparser/gqlerror"
	"go.xixo.com/api/gateway/auth"
	"go.xixo.com/api/gateway/graph/generated"
	"go.xixo.com/api/gateway/graph/marshaller"
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/api/gateway/grpcerror"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"
)

func (r *queryResolver) Roles(ctx context.Context, first int, after *string) (*model.RolesConnection, error) {
	res, err := r.rolesClient.ListRoles(ctx, &identitypb.ListRolesRequest{
		PageSize:  int32(first),
		PageToken: str.Dereference(after),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	edges, err := marshaller.PbToRoleEdges(res.Roles)
	if err != nil {
		return nil, err
	}
	return &model.RolesConnection{
		TotalCount: 0,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage: res.NextPageToken != "",
		},
	}, nil
}

func (r *queryResolver) Role(ctx context.Context, id string) (*model.Role, error) {
	role, err := r.rolesClient.GetRole(ctx, &identitypb.GetRoleRequest{
		Name: roles.Name{RoleID: id}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToRole(role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	user, err := r.usersClient.GetUser(ctx, &identitypb.GetUserRequest{
		Name: users.Name{AccountID: claims.AccountID, UserID: claims.Subject}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Users(ctx context.Context, first int, after *string) (*model.UsersConnection, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	res, err := r.usersClient.ListUsers(ctx, &identitypb.ListUsersRequest{
		Parent:    accounts.Name{AccountID: claims.AccountID}.String(),
		PageSize:  int32(first),
		PageToken: str.Dereference(after),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	edges, err := marshaller.PbToUserEdges(res.Users)
	if err != nil {
		return nil, err
	}
	return &model.UsersConnection{
		TotalCount: 0,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			HasNextPage: res.NextPageToken != "",
		},
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	user, err := r.usersClient.GetUser(ctx, &identitypb.GetUserRequest{
		Name: users.Name{AccountID: claims.AccountID, UserID: id}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Account(ctx context.Context) (*model.Account, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	account, err := r.accountsClient.GetAccount(ctx, &accountpb.GetAccountRequest{
		Name: accounts.Name{AccountID: claims.AccountID}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToAccount(account)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
