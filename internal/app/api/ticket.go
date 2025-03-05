package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type TicketApi interface {
	// GetTicketById returns the ticket with the specified id.
	//
	// May return domain.ErrNoTicket.
	GetTicketById(ctx context.Context, id string) (*domain.Ticket, error)

	// GetTicketByChatId returns an open ticket with the specified ticketId.
	//
	// May return domain.ErrNoTicket.
	GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error)

	// GetTickets returns a slice of tickets with different closed values.
	GetTickets(ctx context.Context, page, perPage int) ([]*domain.Ticket, error)

	// GetTicketsByClosed returns a slice of tickets with the specified closed value.
	GetTicketsByClosed(ctx context.Context, closed bool, page, perPage int) ([]*domain.Ticket, error)

	// OpenTicket opens a ticket and returns it.
	//
	// Returns domain.ErrTicketOpen if there is already an open ticket with the specified chatId.
	//
	// May return domain.ErrTicketOpen.
	OpenTicket(ctx context.Context, chatId int64, topic string) (*domain.Ticket, error)

	// CloseTicketById closes an open ticket with the specified id.
	//
	// May return domain.NoTicket, domain.ErrTicketClosed
	CloseTicketById(ctx context.Context, id string) error

	// CloseTicketByChatId closes an open ticket with the specified chatId.
	//
	// May return domain.ErrTicketClosed
	CloseTicketByChatId(ctx context.Context, chatId int64) error
}
