package services

import (
	"context"
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
	"github.com/rs/zerolog"
	"time"
)

type TicketService struct {
	ticketRepo repo.TicketReadWriter
	log        *zerolog.Logger
}

func NewTicketService(ticketRepo repo.TicketReadWriter, log *zerolog.Logger) api.TicketApi {
	return &TicketService{
		ticketRepo: ticketRepo,
		log:        log,
	}
}

func (service *TicketService) GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error) {
	tickets, err := service.ticketRepo.GetTickets(ctx, limit, offset)
	if err != nil {
		service.log.Err(err).Send()
		return nil, err
	}

	return tickets, nil
}

func (service *TicketService) GetTicketById(ctx context.Context, id string, closed bool) (*domain.Ticket, error) {
	ticket, err := service.ticketRepo.GetTicketById(ctx, id, closed)
	if err != nil {
		service.log.Err(err).Send()
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) GetTicketByChatId(ctx context.Context, chatId int64, closed bool) (*domain.Ticket, error) {
	ticket, err := service.ticketRepo.GetTicketByChatId(ctx, chatId, closed)
	if err != nil {
		service.log.Err(err).Send()
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) OpenTicket(ctx context.Context, chatId int64, topic string) (*domain.Ticket, error) {
	_, err := service.ticketRepo.GetTicketByChatId(ctx, chatId, false)

	if err != nil && !errors.Is(err, domain.ErrTicketNotFound) {
		service.log.Err(err).Send()
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
	ticket, err := service.ticketRepo.GetTicketById(ctx, id, false)

	if err != nil && errors.Is(err, domain.ErrTicketNotFound) {
		return domain.ErrTicketNotFound
	}

	if err != nil {
		return err
	}

	err = service.ticketRepo.SetTicketClosed(ctx, ticket.Id, true)
	if err != nil {
		service.log.Err(err).Send()
		return err
	}

	return nil
}

func (service *TicketService) CloseTicketByChatId(ctx context.Context, chatId int64) error {
	ticket, err := service.ticketRepo.GetTicketByChatId(ctx, chatId, false)

	if err != nil && errors.Is(err, domain.ErrTicketNotFound) {
		return domain.ErrTicketNotFound
	}

	if err != nil {
		service.log.Err(err).Send()
		return err
	}

	if ticket.Closed {
		return domain.ErrTicketAlreadyClosed
	}

	err = service.ticketRepo.SetTicketClosed(ctx, ticket.Id, true)
	if err != nil {
		service.log.Err(err).Send()
		return err
	}

	return nil
}
