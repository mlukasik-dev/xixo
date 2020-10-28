package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"
)

// Resolver .
type Resolver struct {
	authClient     identitypb.AuthClient
	rolesClient    identitypb.RolesClient
	usersClient    identitypb.UsersClient
	accountsClient accountpb.AccountsClient
}

// Clients config struct
type Clients struct {
	AuthClient     identitypb.AuthClient
	RolesClient    identitypb.RolesClient
	UsersClient    identitypb.UsersClient
	AccountsClient accountpb.AccountsClient
}

// NewResolver .
func NewResolver(c *Clients) *Resolver {
	return &Resolver{
		authClient:     c.AuthClient,
		rolesClient:    c.RolesClient,
		usersClient:    c.UsersClient,
		accountsClient: c.AccountsClient,
	}
}
