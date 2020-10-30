//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"
)

// Resolver .
type Resolver struct {
	accountSvcClient  accountpb.AccountServiceClient
	identitySvcClient identitypb.IdentityServiceClient
}

// Clients config struct
type Clients struct {
	AccountSvcClient  accountpb.AccountServiceClient
	IdentitySvcClient identitypb.IdentityServiceClient
}

// NewResolver .
func NewResolver(c *Clients) *Resolver {
	return &Resolver{
		accountSvcClient:  c.AccountSvcClient,
		identitySvcClient: c.IdentitySvcClient,
	}
}
