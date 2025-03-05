package repo

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type TicketReader interface {
	// GetTicketById returns the ticket with the specified id.
	//
	// May return domain.ErrNoTicket
	GetTicketById(ctx context.Context, id string) (*domain.Ticket, error)

	// GetTicketByChatId returns an open ticket with the specified chat id.
	//
	// May return domain.ErrNoTicket
	GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error)

	// GetTickets returns a slice of tickets with different closed values.
	GetTickets(ctx context.Context, page, perPage int) ([]*domain.Ticket, error)

	// GetTicketsByClosed returns a slice of tickets with the specified closed value.
	GetTicketsByClosed(ctx context.Context, closed bool, page, perPage int) ([]*domain.Ticket, error)
}

type TicketWriter interface {
	// AddTicket adds a new ticket
	//
	// May return domain.ErrTicketExists
	AddTicket(ctx context.Context, ticket *domain.Ticket) error

	// SetTicketClosed sets the closed value of a ticket
	//
	// May return domain.ErrNoTicket
	SetTicketClosed(ctx context.Context, id string, closed bool) error
}

type TicketReadWriter interface {
	TicketReader
	TicketWriter
}
