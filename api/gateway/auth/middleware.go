package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/gqlerror"
	"go.xixo.com/api/pkg/token"
)

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if accessToken == "" {
				next.ServeHTTP(w, r)
				return
			}
			splitToken := strings.Split(accessToken, "Bearer ")
			if len(splitToken) != 2 {
				http.Error(w, `{"message": "Invalid token"}`, http.StatusForbidden)
				return
			}
			accessToken = splitToken[1]

			claims, err := verify(accessToken)
			if err != nil {
				http.Error(w, gqlerror.Errorf(err.Error()).Error(), http.StatusForbidden)
				return
			}
			// put it in context
			ctx := token.SetContext(r.Context(), accessToken, claims)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func verify(accessToken string) (*token.JWTClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(
		accessToken,
		&token.JWTClaims{},
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

	claims, ok := parsedToken.Claims.(*token.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
