package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type TicketApi interface {
	// GetTicketById
	//
	// May return domain.ErrNoTicket
	GetTicketById(ctx context.Context, id string) (*domain.Ticket, error)

	// GetTicketByChatId
	//
	// May return domain.ErrNoTicket
	GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error)

	// OpenTicket
	//
	// May return domain.ErrTicketAlreadyOpen
	OpenTicket(ctx context.Context, chatId int64, topic string) error

	// CloseTicketById closes an open ticket with the specified id
	//
	// May return domain.NoTicket, domain.ErrTicketAlreadyClosed
	CloseTicketById(ctx context.Context, id string) error

	// CloseTicketByChatId closes an open ticket with the specified chatId
	//
	// May return domain.ErrTicketAlreadyClosed
	CloseTicketByChatId(ctx context.Context, id string) error
}
