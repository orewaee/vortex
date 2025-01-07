package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type TicketApi interface {
	GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error)
	GetTicketById(ctx context.Context, id string) (*domain.Ticket, error)
	GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error)

	OpenTicket(ctx context.Context, chatId int64) (*domain.Ticket, error)
	CloseTicketById(ctx context.Context, id string) error
	CloseTicketByChatId(ctx context.Context, chatId int64) error
}
