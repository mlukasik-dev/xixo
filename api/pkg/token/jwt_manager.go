package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidToken .
	ErrInvalidToken = errors.Errorf("invalid token")
	// ErrInvalidClaims .
	ErrInvalidClaims = errors.Errorf("invalid claims")
	// ErrUnexpectedTokenSigningMethod .
	ErrUnexpectedTokenSigningMethod = errors.Errorf("unexpected token signing method")
)

// JWTClaims .
type JWTClaims struct {
	jwt.StandardClaims
	Admin     bool        `json:"admin"`
	AccountID *uuid.UUID  `json:"accountId,omitempty"`
	RoleIDs   []uuid.UUID `json:"roleIds"`
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
// accountID is optional
func (manager *JWTManager) Generate(admin bool, accountID *uuid.UUID, id uuid.UUID, roleIDs []uuid.UUID) (string, error) {
	claims := JWTClaims{
		Admin:     admin,
		AccountID: accountID,
		RoleIDs:   roleIDs,
		StandardClaims: jwt.StandardClaims{
			Subject:   id.String(),
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
