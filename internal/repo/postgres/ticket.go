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

func NewTicketRepo(pool *pgxpool.Pool) repo.TicketReadWriter {
	return &TicketRepo{pool}
}

func scanTicket(row pgx.CollectableRow) (*domain.Ticket, error) {
	ticket := new(domain.Ticket)
	err := row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.Closed, &ticket.CreatedAt)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) GetTicketById(ctx context.Context, id string) (*domain.Ticket, error) {
	var ticket *domain.Ticket
	err := withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, "SELECT * FROM tickets WHERE id = $1", id)
		return row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.Closed, &ticket.CreatedAt)
	})

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoTicket
	}

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) GetTicketByChatId(ctx context.Context, chatId int64) (*domain.Ticket, error) {
	var ticket *domain.Ticket
	err := withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, "SELECT * FROM tickets WHERE chat_id = $1 AND closed = false", chatId)
		return row.Scan(&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.Closed, &ticket.CreatedAt)
	})

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoTicket
	}

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *TicketRepo) GetTickets(ctx context.Context, page, perPage int) ([]*domain.Ticket, error) {
	var tickets []*domain.Ticket
	err := withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		offset := page * perPage
		rows, err := tx.Query(ctx, "SELECT * FROM tickets LIMIT $1 OFFSET $2", perPage, offset)
		if err != nil {
			return err
		}

		tickets, err = pgx.CollectRows[*domain.Ticket](rows, scanTicket)
		return err
	})

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (repo *TicketRepo) GetTicketsByClosed(ctx context.Context, closed bool, page, perPage int) ([]*domain.Ticket, error) {
	var tickets []*domain.Ticket
	err := withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		offset := page * perPage
		rows, err := tx.Query(ctx, "SELECT * FROM tickets WHERE closed = $1 LIMIT $2 OFFSET $3", closed, perPage, offset)
		if err != nil {
			return err
		}

		tickets, err = pgx.CollectRows[*domain.Ticket](rows, scanTicket)
		return err
	})

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (repo *TicketRepo) AddTicket(ctx context.Context, ticket *domain.Ticket) error {
	return withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		exists := false
		err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 as alias FROM tickets WHERE id = $1) as E", ticket.Id).Scan(&exists)
		if err != nil {
			return err
		}

		if exists {
			return domain.ErrTicketExists
		}

		_, err = tx.Exec(ctx, "INSERT INTO tickets (id, chat_id, topic, closed, created_at) VALUES ($1, $2, $3, $4, $5)",
			&ticket.Id, &ticket.ChatId, &ticket.Topic, &ticket.Closed, &ticket.CreatedAt)
		return err
	})
}

func (repo *TicketRepo) SetTicketClosedById(ctx context.Context, id string, closed bool) error {
	return withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		exists := false
		err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 as alias FROM tickets WHERE id = $1) as E", id).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			return domain.ErrNoTicket
		}

		_, err = tx.Exec(ctx, "UPDATE tickets SET closed = $1 WHERE id = $2", closed, id)
		return err
	})
}
