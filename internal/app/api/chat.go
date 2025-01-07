package api

import "github.com/orewaee/vortex/internal/app/domain"

type ChatApi interface {
	Connect() chan *domain.Message
	Disconnect(connection chan *domain.Message)
	SendMessage(sender string, fromSupport bool, ticketId string, text string) error
}
