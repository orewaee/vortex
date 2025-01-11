package api

import "context"

type AuthApi interface {
	Login(ctx context.Context, name string, password string) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
}
