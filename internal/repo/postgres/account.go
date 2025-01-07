package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
)

type AccountRepo struct {
	pool *pgxpool.Pool
}

func NewAccountRepo(pool *pgxpool.Pool) repo.AccountReadWriter {
	return &AccountRepo{pool}
}

func (repo *AccountRepo) GetAccountById(ctx context.Context, id string) (*domain.Account, error) {
	panic("unimplemented")
}

func (repo *AccountRepo) GetAccountByName(ctx context.Context, name string) (*domain.Account, error) {
	panic("unimplemented")
}

func (repo *AccountRepo) AddAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	panic("unimplemented")
}

func (repo *AccountRepo) RemoveAccountById(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (repo *AccountRepo) RemoveAccountByName(ctx context.Context, name string) error {
	panic("unimplemented")
}
