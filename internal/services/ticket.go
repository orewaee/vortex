package services

import (
	"context"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"github.com/orewaee/vortex/internal/utils"
	"github.com/rs/zerolog"
	"time"
)

type TicketService struct {
	ticketRepo repo.TicketReadWriter
	log        *zerolog.Logger
}

func NewTicketService(
	ticketRepo repo.TicketReadWriter,
	log *zerolog.Logger) api.TicketApi {
	return &TicketService{
		ticketRepo: ticketRepo,
		log:        log,
	}
}

func (service *TicketService) GetTicketById(ctx context.Context, id string) (*domain.Ticket, error) {
	ticket, err := service.ticketRepo.GetTicketById(ctx, id)
	if err != nil {
		service.log.Error().Err(err).Send()
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error) {
	ticket, err := service.ticketRepo.GetTicketByChatId(ctx, chatId)
	if err != nil {
		service.log.Error().Err(err).Send()
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) GetTickets(ctx context.Context, page, perPage int) ([]*domain.Ticket, error) {
	tickets, err := service.ticketRepo.GetTickets(ctx, page, perPage)
	if err != nil {
		service.log.Error().Err(err).Send()
		return nil, err
	}

	return tickets, nil
}

func (service *TicketService) GetTicketsByClosed(ctx context.Context, closed bool, page, perPage int) ([]*domain.Ticket, error) {
	tickets, err := service.ticketRepo.GetTicketsByClosed(ctx, closed, page, perPage)
	if err != nil {
		service.log.Error().Err(err).Send()
		return nil, err
	}

	return tickets, nil
}

func (service *TicketService) OpenTicket(ctx context.Context, chatId int64, topic string) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		Id:        utils.MustNewId(),
		ChatId:    chatId,
		Topic:     topic,
		Closed:    false,
		CreatedAt: time.Now(),
	}

	if err := service.ticketRepo.AddTicket(ctx, ticket); err != nil {
		service.log.Error().Err(err).Send()
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) CloseTicketById(ctx context.Context, id string) error {
	err := service.ticketRepo.SetTicketClosed(ctx, id, true)
	if err != nil {
		service.log.Error().Err(err).Send()
		return err
	}

	return nil
}

func (service *TicketService) CloseTicketByChatId(ctx context.Context, chatId int64) error {
	ticket, err := service.GetTicketByChatId(ctx, chatId)
	if err != nil {
		service.log.Error().Err(err).Send()
		return err
	}

	err = service.ticketRepo.SetTicketClosed(ctx, ticket.Id, false)
	if err != nil {
		service.log.Error().Err(err).Send()
		return err
	}

	return nil
}
