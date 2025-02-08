package duckdb

import (
	"context"
	"database/sql"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
)

type TicketRepo struct {
	db *sql.DB
}

func NewTicketRepo(db *sql.DB) repo.TicketReadWriter {
	return &TicketRepo{db: db}
}

func (repo *TicketRepo) GetTicketById(ctx context.Context, id string) (*domain.Ticket, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	ticket := new(domain.Ticket)
	err = tx.QueryRowContext(ctx, "SELECT * FROM tickets WHERE id = ?", id).
		Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.Closed, &ticket.CreatedAt)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TicketRepo) GetTickets(ctx context.Context, page, perPage int) ([]*domain.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TicketRepo) GetTicketsByClosed(ctx context.Context, closed bool, page, perPage int) ([]*domain.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TicketRepo) AddTicket(ctx context.Context, ticket *domain.Ticket) error {
	//TODO implement me
	panic("implement me")
}

func (repo *TicketRepo) SetTicketClosed(ctx context.Context, id string, closed bool) error {
	//TODO implement me
	panic("implement me")
}
