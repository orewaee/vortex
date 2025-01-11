package repo

import (
	"context"
    "github.com/orewaee/vortex/internal/app/domain"
)

type ChatReader interface {
	GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error)
}

type ChatWriter interface {
    AddMessage(ctx context.Context, message *domain.Message) error
    RemoveMessageById(ctx context.Context, id string) error
}

type ChatReadWriter interface {
    ChatReader
    ChatWriter
}

