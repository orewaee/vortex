package api

import (
	"context"
	"github.com/orewaee/vortex/internal/app/domain"
)

type ChatApi interface {
	Subscribe() chan *domain.Message
	Unsubscribe(conn chan *domain.Message)

	GetMessageHistory(ctx context.Context, ticketId string, page, perPage int) ([]*domain.Message, error)
	SendMessage(ctx context.Context, sender string, fromSupport bool, ticketId string, text string) error
}
