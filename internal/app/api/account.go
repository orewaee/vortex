package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type AccountApi interface {
	GetAccountById(ctx context.Context, id string) (*domain.Account, error)
	GetAccountByName(ctx context.Context, name string) (*domain.Account, error)

	AddAccount(ctx context.Context, name, password string) (*domain.Account, error)
	RemoveAccountById(ctx context.Context, id string) error
	RemoveAccountByName(ctx context.Context, name string) error
}
