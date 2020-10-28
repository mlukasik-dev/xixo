package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/gqlerror"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var claimsCtxKey = &contextKey{"claims"}
var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}
			splitToken := strings.Split(token, "Bearer ")
			if len(splitToken) != 2 {
				http.Error(w, `{"message": "Invalid token"}`, http.StatusForbidden)
				return
			}
			token = splitToken[1]

			claims, err := verify(token)
			if err != nil {
				http.Error(w, gqlerror.Errorf(err.Error()).Error(), http.StatusForbidden)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), tokenCtxKey, token)
			ctx = context.WithValue(ctx, claimsCtxKey, claims)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ClaimsFromContext finds the user from the context. REQUIRES Middleware to have run.
func ClaimsFromContext(ctx context.Context) (*JWTClaims, bool) {
	claims, ok := ctx.Value(claimsCtxKey).(*JWTClaims)
	return claims, ok
}

// TokenFromContext finds the user from the context. REQUIRES Middleware to have run.
func TokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenCtxKey).(string)
	return token, ok
}

// JWTClaims .
type JWTClaims struct {
	jwt.StandardClaims
	Admin     bool     `json:"admin"`
	AccountID string   `json:"accountId"`
	Roles     []string `json:"roles"`
}

func verify(accessToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			// TODO: get from .env
			return []byte("secret"), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
