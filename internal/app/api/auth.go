package api

import "context"

type AuthApi interface {
	Login(ctx context.Context, name string, password string) (string, string, error)
	Refresh(ctx context.Context, token string) (string, string, error)
}
