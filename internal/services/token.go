package services

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"time"
)

type JwtTokenService struct {
	tokenRepo repo.TokenReadWriter
}

func NewJwtTokenService(tokenRepo repo.TokenReadWriter) api.TokenApi {
	return &JwtTokenService{tokenRepo}
}

func (service *JwtTokenService) CreateToken(claims map[string]interface{}, key string) (string, error) {
	var mapClaims jwt.MapClaims = claims

	unsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	signed, err := unsigned.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (service *JwtTokenService) WhitelistToken(ctx context.Context, token string, lifetime time.Duration) error {
	return service.tokenRepo.AddToken(ctx, token, lifetime)
}

func (service *JwtTokenService) GetTokenClaims(token string, key string) (map[string]interface{}, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, domain.ErrInvalidToken
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.ErrMissingClaims
	}

	return claims, nil
}

func (service *JwtTokenService) RefreshToken(ctx context.Context, token string) (string, string, error) {
	exists, err := service.tokenRepo.TokenExists(ctx, token)

	if err != nil || !exists {
		return "", "", domain.ErrInvalidToken
	}

	if err := service.tokenRepo.RemoveToken(ctx, token); err != nil {
		return "", "", err
	}

	mapClaims, err := service.GetTokenClaims(token, typedenv.String("REFRESH_KEY"))

	if err != nil {
		return "", "", domain.ErrInvalidToken
	}

	now := time.Now()

	access, err := service.CreateToken(map[string]interface{}{
		"iss":   "vortex",
		"name":  mapClaims["name"],
		"perms": mapClaims["perms"],
		"exp":   now.Add(typedenv.Duration("ACCESS_LIFETIME")).Unix(),
		"iat":   now.Unix(),
	}, typedenv.String("ACCESS_KEY"))

	if err != nil {
		return "", "", err
	}

	lifetime := typedenv.Duration("REFRESH_LIFETIME")
	refresh, err := service.CreateToken(map[string]interface{}{
		"iss":   "vortex",
		"name":  mapClaims["name"],
		"perms": mapClaims["perms"],
		"exp":   now.Add(lifetime).Unix(),
		"iat":   now.Unix(),
	}, typedenv.String("REFRESH_KEY"))

	if err != nil {
		return "", "", err
	}

	if err := service.tokenRepo.AddToken(ctx, refresh, lifetime); err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
