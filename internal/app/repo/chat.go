package repo

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

// ChatReader contains methods for reading chat messages
type ChatReader interface {
	GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error)
}

// ChatWriter contains methods for writing chat messages
type ChatWriter interface {
	AddMessage(ctx context.Context, message *domain.Message) error
	RemoveMessageById(ctx context.Context, id string) error
}

// ChatReadWriter is a wrapper for ChatReader and ChatWriter
type ChatReadWriter interface {
	ChatReader
	ChatWriter
}
