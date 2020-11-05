package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sync"

	"github.com/vektah/gqlparser/gqlerror"
	"go.xixo.com/api/gateway/graph/generated"
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/api/gateway/graph/transform"
	"go.xixo.com/api/gateway/grpcerror"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/protobuf/identitypb"

	"github.com/google/uuid"
)

func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]*model.Role, error) {
	roleSlice := make([]*model.Role, len(obj.Roles))
	wg := sync.WaitGroup{}
	wg.Add(len(obj.Roles))
	errCh := make(chan error)
	waitCh := make(chan struct{})

	go func() {
		for i, roleID := range obj.Roles {
			go func(i int, id string) {
				defer wg.Done()
				roleID, err := uuid.Parse(id)
				if err != nil {
					errCh <- gqlerror.Errorf("invalid role id")
					return
				}
				rolePb, err := r.identitySvcClient.GetRole(ctx, &identitypb.GetRoleRequest{
					Name: roles.Name{RoleID: roleID}.String(),
				})
				if err != nil {
					errCh <- grpcerror.GetError(err)
					return
				}
				role, err := transform.PbToRole(rolePb)
				if err != nil {
					errCh <- err
					return
				}
				roleSlice[i] = role
			}(i, roleID)
		}
		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		return roleSlice, nil
	case err := <-errCh:
		return nil, err
	}
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
