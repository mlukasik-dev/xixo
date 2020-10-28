package config

import (
	"os"
)

// JWTConfig jwt configuration object
type JWTConfig struct {
	Secret string
}

// JWT stores configuration for jwt middleware
var JWT *JWTConfig

func init() {
	JWT = &JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}
