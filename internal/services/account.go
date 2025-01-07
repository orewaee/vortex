package services

import (
	"context"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"time"
)

type AccountService struct {
}

func NewAccountService() api.AccountApi {
	return &AccountService{}
}

func (service AccountService) GetAccountById(ctx context.Context, id string) (*domain.Account, error) {
	superId := typedenv.String("SUPER_ID")

	if id != superId {
		return nil, domain.ErrAccountNotFound
	}

	return &domain.Account{
		Id:        superId,
		Name:      typedenv.String("SUPER_NAME"),
		Password:  typedenv.String("SUPER_PASSWORD"),
		Perms:     domain.PermSuper,
		CreatedAt: time.Now(),
	}, nil
}

func (service AccountService) GetAccountByName(ctx context.Context, name string) (*domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (service AccountService) AddAccount(ctx context.Context, name, password string) (*domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (service AccountService) RemoveAccountById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (service AccountService) RemoveAccountByName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
