package services

import (
	"context"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"time"
)

type AuthService struct {
	tokenService api.TokenApi
}

func NewAuthService(tokenService api.TokenApi) api.AuthApi {
	return &AuthService{tokenService}
}

func (service *AuthService) Login(ctx context.Context, name string, password string) (string, string, error) {
	isSuper := name == typedenv.String("SUPER_NAME") &&
		password == typedenv.String("SUPER_PASSWORD")

	if !isSuper {
		return "", "", domain.ErrInvalidCredentials
	}

	now := time.Now()

	access, err := service.tokenService.CreateToken(map[string]interface{}{
		"iss":   "vortex",
		"name":  name,
		"perms": domain.PermSuper,
		"exp":   now.Add(typedenv.Duration("ACCESS_LIFETIME")).Unix(),
		"iat":   now.Unix(),
	}, typedenv.String("ACCESS_KEY"))

	if err != nil {
		return "", "", err
	}

	lifetime := typedenv.Duration("REFRESH_LIFETIME")
	refresh, err := service.tokenService.CreateToken(map[string]interface{}{
		"iss":   "vortex",
		"name":  name,
		"perms": domain.PermSuper,
		"exp":   now.Add(lifetime).Unix(),
		"iat":   now.Unix(),
	}, typedenv.String("REFRESH_KEY"))

	if err != nil {
		return "", "", err
	}

	if err := service.tokenService.WhitelistToken(ctx, refresh, lifetime); err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (service *AuthService) Refresh(ctx context.Context, token string) (string, string, error) {
	return service.tokenService.RefreshToken(ctx, token)
}
