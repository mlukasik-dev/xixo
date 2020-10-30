package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/vektah/gqlparser/gqlerror"
	"go.xixo.com/api/gateway/graph/generated"
	"go.xixo.com/api/gateway/graph/marshaller"
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/api/gateway/grpcerror"
	"go.xixo.com/api/pkg/token"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/protobuf/identitypb"
)

func (r *mutationResolver) Login(ctx context.Context, accountID string, email string, password string) (*model.LoginPayload, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	res, err := r.identitySvcClient.Login(ctx, &identitypb.LoginRequest{
		AccountId: accountID,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	return &model.LoginPayload{Token: res.AccessToken}, nil
}

func (r *mutationResolver) Register(ctx context.Context, accountID string, email string, password string) (*model.RegisterPayload, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	res, err := r.identitySvcClient.Register(ctx, &identitypb.RegisterRequest{
		AccountId: accountID,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	return &model.RegisterPayload{Token: res.AccessToken}, nil
}

func (r *mutationResolver) CreateRole(ctx context.Context, input *model.CreateRoleInput) (*model.Role, error) {
	role, err := r.identitySvcClient.CreateRole(ctx, marshaller.CreateRoleInputToPB(input))
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToRole(role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input *model.UpdateRoleInput) (*model.Role, error) {
	req, err := marshaller.UpdateRoleInputToPB(id, input)
	if err != nil {
		return nil, err
	}
	role, err := r.identitySvcClient.UpdateRole(ctx, req)
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToRole(role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*model.Role, error) {
	_, err := r.identitySvcClient.DeleteRole(ctx, &identitypb.DeleteRoleRequest{
		Name: roles.Name{RoleID: id}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	claims, ok := token.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	user, err := r.identitySvcClient.CreateUser(ctx, marshaller.CreateUserInputToPB(claims.AccountID, input))
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UpdateUserInput) (*model.User, error) {
	claims, ok := token.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	req, err := marshaller.UpdateUserInputToPB(claims.AccountID, id, input)
	if err != nil {
		return nil, err
	}
	user, err := r.identitySvcClient.UpdateUser(ctx, req)
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	res, err := marshaller.PbToUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserPayload, error) {
	claims, ok := token.ClaimsFromContext(ctx)
	if !ok {
		return nil, gqlerror.Errorf("cannot get claims")
	}
	_, err := r.identitySvcClient.DeleteUser(ctx, &identitypb.DeleteUserRequest{
		Name: users.Name{AccountID: claims.AccountID, UserID: id}.String(),
	})
	if err != nil {
		return nil, grpcerror.GetError(err)
	}
	return &model.DeleteUserPayload{}, nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.UpdateAccountInput) (*model.Account, error) {
	return nil, &gqlerror.Error{}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
