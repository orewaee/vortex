package postgres

import (
	"context"

	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
)

type ChatRepo struct {
	pool *pgxpool.Pool
}

func NewChatRepo(pool *pgxpool.Pool) (repo.ChatReadWriter, error) {
	sql := `
	    CREATE TABLE IF NOT EXISTS messages (
	    id CHAR(8) PRIMARY KEY,
        sender VARCHAR(32),
        from_support BOOLEAN,
        ticket_id CHAR(8) REFERENCES tickets(id),
        text TEXT,
	    created_at TIMESTAMP WITH TIME ZONE
	)
	`

	if _, err := pool.Exec(context.Background(), sql); err != nil {
		return nil, err
	}

	return &ChatRepo{pool}, nil
}

func (repo *ChatRepo) GetMessageHistory(ctx context.Context, ticketId string, limit, offset int) ([]*domain.Message, error) {
	rows, err := repo.pool.Query(ctx, "SELECT * FROM messages WHERE ticket_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3", ticketId, limit, offset)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoMessages
	}

	if err != nil {
		return nil, err
	}

	messages, err := pgx.CollectRows[*domain.Message](rows, func(row pgx.CollectableRow) (*domain.Message, error) {
		message := new(domain.Message)
		if err := row.Scan(&message.Id, &message.Sender, &message.FromSupport, &message.TicketId, &message.Text, &message.CreatedAt); err != nil {
			return nil, err
		}

		return message, nil
	})

	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, domain.ErrNoMessages
	}

	return messages, nil

}

func (repo *ChatRepo) AddMessage(ctx context.Context, message *domain.Message) error {
	sql := "INSERT INTO messages (id, sender, from_support, ticket_id, text, created_at) values ($1, $2, $3, $4, $5, $6)"
	_, err := repo.pool.Exec(ctx, sql, message.Id, message.Sender, message.FromSupport, message.TicketId, message.Text, message.CreatedAt)

	return err
}

func (repo *ChatRepo) RemoveMessageById(ctx context.Context, id string) error {
	sql := "DELETE FROM messages WHERE id = $1"
	_, err := repo.pool.Exec(ctx, sql, id)

	return err
}
