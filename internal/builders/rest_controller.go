package builders

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/driving"
	"github.com/rs/zerolog"
)

type RestControllerBuilder interface {
	Builder[driving.Controller]

	// AccountRepo(accountRepo repo.AccountReadWriter) RestControllerBuilder
	// LoginCodeRepo(loginCodeRepo repo.LoginCo) RestControllerBuilder
	// TokenRepo(tokenRepo repo.TokenReadWriter) RestControllerBuilder
}

type restControllerBuilder struct {
	authApi   api.AuthApi
	tokenApi  api.TokenApi
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	logger    *zerolog.Logger
}

/*
func NewRestControllerBuilder() RestControllerBuilder {
	return restControllerBuilder{}
}
*/
