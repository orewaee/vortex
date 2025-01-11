package services

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"github.com/orewaee/vortex/internal/broker"
	"time"
)

type ChatService struct {
	broker   *broker.Broker[*domain.Message]
	chatRepo repo.ChatReadWriter
}

func NewChatService(chatRepo repo.ChatReadWriter) api.ChatApi {
	return &ChatService{
		broker:   broker.New[*domain.Message](),
		chatRepo: chatRepo,
	}
}

func (service *ChatService) Subscribe() chan *domain.Message {
	return service.broker.Subscribe()
}

func (service *ChatService) Unsubscribe(connection chan *domain.Message) {
	service.broker.Unsubscribe(connection)
}

func (service *ChatService) GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error) {
	return service.chatRepo.GetMessageHistory(ctx, ticketId, limit, offset)
}

func (service *ChatService) SendMessage(ctx context.Context, sender string, fromSupport bool, ticketId string, text string) error {
	id := gonanoid.MustGenerate(typedenv.String("ALPHABET"), 8)

	message := &domain.Message{
		Id:          id,
		Sender:      sender,
		FromSupport: fromSupport,
		TicketId:    ticketId,
		Text:        text,
		CreatedAt:   time.Now(),
	}

	service.broker.Publish(message)

	return service.chatRepo.AddMessage(ctx, message)
}
