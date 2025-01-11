package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

// ChatApi contains methods for working with chat
type ChatApi interface {
	Subscribe() chan *domain.Message
	Unsubscribe(connection chan *domain.Message)

	GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error)
	SendMessage(ctx context.Context, sender string, fromSupport bool, ticketId string, text string) error
}
