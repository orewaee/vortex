package api

import (
	"context"
	"time"
)

// TokenApi contains methods for working with tokens
type TokenApi interface {
	CreateToken(claims map[string]interface{}, key string) (string, error)
	WhitelistToken(ctx context.Context, token string, lifetime time.Duration) error
	GetTokenClaims(token string, key string) (map[string]interface{}, error)
	RefreshToken(ctx context.Context, token string) (string, string, error)
}
