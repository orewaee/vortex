package services

import (
	"context"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"github.com/orewaee/vortex/internal/broker"
	"github.com/orewaee/vortex/internal/utils"
	"github.com/rs/zerolog"
	"time"
)

type ChatService struct {
	broker   *broker.Broker[*domain.Message]
	chatRepo repo.ChatReadWriter
	log      *zerolog.Logger
}

func NewChatService(chatRepo repo.ChatReadWriter, log *zerolog.Logger) api.ChatApi {
	return &ChatService{
		broker:   broker.New[*domain.Message](),
		chatRepo: chatRepo,
		log:      log,
	}
}

func (service *ChatService) Subscribe() chan *domain.Message {
	return service.broker.Subscribe()
}

func (service *ChatService) Unsubscribe(connection chan *domain.Message) {
	service.broker.Unsubscribe(connection)
}

func (service *ChatService) GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error) {
	messages, err := service.chatRepo.GetMessageHistory(ctx, ticketId, limit, offset)
	if err != nil {
		service.log.Err(err).Send()
		return nil, err
	}

	return messages, nil
}

func (service *ChatService) SendMessage(ctx context.Context, sender string, fromSupport bool, ticketId string, text string) error {
	message := &domain.Message{
		Id:          utils.MustNewId(),
		Sender:      sender,
		FromSupport: fromSupport,
		TicketId:    ticketId,
		Text:        text,
		CreatedAt:   time.Now(),
	}

	service.broker.Publish(message)

	err := service.chatRepo.AddMessage(ctx, message)
	if err != nil {
		service.log.Err(err).Send()
		return err
	}

	return nil
}
