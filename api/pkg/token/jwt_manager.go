package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// ErrInvalidToken .
var ErrInvalidToken = errors.Errorf("invalid token")

// ErrInvalidClaims .
var ErrInvalidClaims = errors.Errorf("invalid claims")

// ErrUnexpectedTokenSigningMethod .
var ErrUnexpectedTokenSigningMethod = errors.Errorf("unexpected token signing method")

// JWTClaims .
type JWTClaims struct {
	jwt.StandardClaims
	Admin     bool     `json:"admin"`
	AccountID string   `json:"accountId"`
	RoleIDs   []string `json:"roleIds"`
}

// JWTManager .
type JWTManager struct {
	secret        string
	tokenDuration time.Duration
}

// NewJWTManager .
func NewJWTManager(secret string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secret, tokenDuration}
}

// Generate .
func (manager *JWTManager) Generate(admin bool, accountID, id string, roleIDs []string) (string, error) {
	claims := JWTClaims{
		Admin:     admin,
		AccountID: accountID,
		RoleIDs:   roleIDs,
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secret))
}

// Verify .
func (manager *JWTManager) Verify(accessToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedTokenSigningMethod
			}

			return []byte(manager.secret), nil
		},
	)

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}
