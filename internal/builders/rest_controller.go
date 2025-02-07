package builders

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/repo"
	"github.com/rs/zerolog"
)

type RestControllerBuilder interface {
	Builder[api.Controller]
	AccountRepo(accountRepo repo.AccountReadWriter) AuthServiceBuilder
	LoginCodeRepo(loginCodeRepo repo.LoginCodeReadWriter) AuthServiceBuilder
	TokenRepo(tokenRepo repo.TokenReadWriter) AuthServiceBuilder
}

type restControllerBuilder struct {
	authApi   api.AuthApi
	tokenApi  api.TokenApi
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	logger    *zerolog.Logger
}

func NewRestControllerBuilder() RestControllerBuilder {
	return restControllerBuilder{}
}
