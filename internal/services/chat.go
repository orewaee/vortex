package services

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/broker"
)

type ChatService struct {
	broker *broker.Broker[*domain.Message]
}

func NewChatService() api.ChatApi {
	return &ChatService{
		broker: broker.New[*domain.Message](),
	}
}

func (service *ChatService) SendMessage(sender string, fromSupport bool, ticketId string, text string) error {
	message := &domain.Message{
		Sender:      sender,
		FromSupport: fromSupport,
		TicketId:    ticketId,
		Text:        text,
	}

	service.broker.Publish(message)

	return nil
}

func (service *ChatService) Connect() chan *domain.Message {
	return service.broker.Subscribe()
}

func (service *ChatService) Disconnect(connection chan *domain.Message) {
	service.broker.Unsubscribe(connection)
}
