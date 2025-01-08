package services

import (
	"context"
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"time"
)

type TicketService struct {
	ticketRepo repo.TicketReadWriter
}

func NewTicketService(ticketRepo repo.TicketReadWriter) api.TicketApi {
	return &TicketService{ticketRepo}
}

func (service *TicketService) GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error) {
	return service.ticketRepo.GetTickets(ctx, limit, offset)
}

func (service *TicketService) GetTicketById(ctx context.Context, id string) (*domain.Ticket, error) {
	return service.ticketRepo.GetTicketById(ctx, id)
}

func (service *TicketService) GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error) {
	return service.ticketRepo.GetTicketByChatId(ctx, chatId)
}

func (service *TicketService) OpenTicket(ctx context.Context, chatId int64, topic string) (*domain.Ticket, error) {
	_, err := service.ticketRepo.GetTicketByChatId(ctx, chatId)

	if err != nil && !errors.Is(err, domain.ErrTicketNotFound) {
		return nil, err
	}

	if err == nil {
		return nil, domain.ErrTicketAlreadyExists
	}

	ticket := &domain.Ticket{
		Id:        gonanoid.MustGenerate(typedenv.String("ALPHABET"), 8),
		ChatId:    chatId,
		Topic:     topic,
		CreatedAt: time.Now(),
	}

	if err := service.ticketRepo.AddTicket(ctx, ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) CloseTicketById(ctx context.Context, id string) error {
	_, err := service.ticketRepo.GetTicketById(ctx, id)

	if err != nil && errors.Is(err, domain.ErrTicketNotFound) {
		return domain.ErrTicketNotFound
	}

	if err != nil {
		return err
	}

	return service.ticketRepo.RemoveTicketById(ctx, id)
}

func (service *TicketService) CloseTicketByChatId(ctx context.Context, chatId int64) error {
	_, err := service.ticketRepo.GetTicketByChatId(ctx, chatId)

	if err != nil && errors.Is(err, domain.ErrTicketNotFound) {
		return domain.ErrTicketNotFound
	}

	if err != nil {
		return err
	}

	return service.ticketRepo.RemoveTicketByChatId(ctx, chatId)
}
