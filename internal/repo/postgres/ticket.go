package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
)

type TicketRepo struct {
	pool *pgxpool.Pool
}

func NewTicketRepo(pool *pgxpool.Pool) (repo.TicketReadWriter, error) {
	sql := `
	CREATE TABLE IF NOT EXISTS tickets (
	    id CHAR(8) PRIMARY KEY,
	    chat_id BIGINT UNIQUE NOT NULL,
	    topic VARCHAR(64) NOT NULL,
	    created_at TIMESTAMP WITH TIME ZONE
	)
	`

	if _, err := pool.Exec(context.Background(), sql); err != nil {
		return nil, err
	}

	return &TicketRepo{pool}, nil
}

func (repo *TicketRepo) GetTickets(ctx context.Context, limit, offset int) ([]*domain.Ticket, error) {
	rows, err := repo.pool.Query(ctx, "SELECT * FROM tickets LIMIT $1 OFFSET $2", limit, offset)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoTickets
	}

	if err != nil {
		return nil, err
	}

	tickets, err := pgx.CollectRows[*domain.Ticket](rows, func(row pgx.CollectableRow) (*domain.Ticket, error) {
		ticket := new(domain.Ticket)
		if err := row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.CreatedAt); err != nil {
			return nil, err
		}

		return ticket, nil
	})

	if err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, domain.ErrNoTickets
	}

	return tickets, nil
}

func (repo *TicketRepo) GetTicketById(ctx context.Context, id string) (*domain.Ticket, error) {
	row := repo.pool.QueryRow(ctx, "SELECT * FROM tickets WHERE id = $1", id)

	ticket := new(domain.Ticket)
	err := row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.CreatedAt)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrTicketNotFound
	}

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error) {
	row := repo.pool.QueryRow(ctx, "SELECT * FROM tickets WHERE chat_id = $1", chatId)

	ticket := new(domain.Ticket)
	err := row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.CreatedAt)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrTicketNotFound
	}

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) AddTicket(ctx context.Context, ticket *domain.Ticket) error {
	sql := "INSERT INTO tickets (id, chat_id, topic, created_at) values ($1, $2, $3, $4)"
	_, err := repo.pool.Exec(ctx, sql, ticket.Id, ticket.ChatId, ticket.Topic, ticket.CreatedAt)

	return err
}

func (repo *TicketRepo) RemoveTicketById(ctx context.Context, id string) error {
	sql := "DELETE FROM tickets WHERE id = $1"
	_, err := repo.pool.Exec(ctx, sql, id)

	return err
}

func (repo *TicketRepo) RemoveTicketByChatId(ctx context.Context, chatId int64) error {
	sql := "DELETE FROM tickets WHERE chat_id = $1"
	_, err := repo.pool.Exec(ctx, sql, chatId)

	return err
}
