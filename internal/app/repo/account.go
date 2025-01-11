package repo

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

// AccountReader contains methods for reading accounts
type AccountReader interface {
	GetAccountById(ctx context.Context, id string) (*domain.Account, error)
	GetAccountByName(ctx context.Context, name string) (*domain.Account, error)
}

// AccountWriter contains methods for writing accounts
type AccountWriter interface {
	AddAccount(ctx context.Context, account *domain.Account) (*domain.Account, error)
	RemoveAccountById(ctx context.Context, id string) error
	RemoveAccountByName(ctx context.Context, name string) error
}

// AccountReadWriter is a wrapper for AccountReader and AccountWriter
type AccountReadWriter interface {
	AccountReader
	AccountWriter
}
