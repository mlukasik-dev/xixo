package token

import "context"

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var claimsCtxKey = &contextKey{"claims"}
var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

// ClaimsFromContext finds the user from the context. REQUIRES Middleware to have run.
func ClaimsFromContext(ctx context.Context) (*JWTClaims, bool) {
	claims, ok := ctx.Value(claimsCtxKey).(*JWTClaims)
	return claims, ok
}

// SetContext .
func SetContext(ctx context.Context, token string, claims *JWTClaims) context.Context {
	return context.WithValue(
		context.WithValue(ctx, tokenCtxKey, token),
		tokenCtxKey, claims,
	)
}

// FromContext finds the user from the context. REQUIRES Middleware to have run.
func FromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenCtxKey).(string)
	return token, ok
}
