package repo

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

// TicketReader contains methods for reading tickets
type TicketReader interface {
	GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error)
	GetTicketById(ctx context.Context, id string) (*domain.Ticket, error)
	GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error)
}

// TicketWriter contains methods for writing tickets
type TicketWriter interface {
	AddTicket(ctx context.Context, ticket *domain.Ticket) error
	RemoveTicketById(ctx context.Context, id string) error
	RemoveTicketByChatId(ctx context.Context, chatId int64) error
}

// TicketReadWriter is a wrapper for TicketReader and TicketWriter
type TicketReadWriter interface {
	TicketReader
	TicketWriter
}
