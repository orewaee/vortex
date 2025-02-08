package repo

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

// TicketReader contains methods for reading tickets
type TicketReader interface {
	// GetTickets returns a slice of tickets (open and closed)
	GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error)

	// GetTicketById returns a ticket with the specified id
	GetTicketById(ctx context.Context, id string, closed bool) (*domain.Ticket, error)

	// GetTicketByChatId returns a ticket with the specified chatId
	GetTicketByChatId(ctx context.Context, chatId int64, closed bool) (*domain.Ticket, error)
}

// TicketWriter contains methods for writing tickets
type TicketWriter interface {
	AddTicket(ctx context.Context, ticket *domain.Ticket) error
	RemoveTicketById(ctx context.Context, id string) error
	RemoveTicketByChatId(ctx context.Context, chatId int64) error
	SetTicketClosed(ctx context.Context, id string, closed bool) error
}

// TicketReadWriter is a wrapper for TicketReader and TicketWriter
type TicketReadWriter interface {
	TicketReader
	TicketWriter
}
