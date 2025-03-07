package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/app/repo"
)

type ChatRepo struct {
	pool *pgxpool.Pool
}

func NewChatRepo(pool *pgxpool.Pool) repo.ChatReadWriter {
	return &ChatRepo{pool}
}

func scanMessage(row pgx.CollectableRow) (*domain.Message, error) {
	message := new(domain.Message)
	err := row.Scan(&message.Id, &message.Sender, &message.FromSupport, &message.TicketId, &message.Text, &message.CreatedAt)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (repo *ChatRepo) GetMessageHistory(ctx context.Context, ticketId string, page, perPage int) ([]*domain.Message, error) {
	messages := make([]*domain.Message, 0)
	err := withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		offset := page * perPage
		rows, err := repo.pool.Query(ctx, "SELECT * FROM messages WHERE ticket_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3", ticketId, perPage, offset)
		if err != nil {
			return err
		}

		messages, err = pgx.CollectRows[*domain.Message](rows, scanMessage)
		return err
	})

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo *ChatRepo) AddMessage(ctx context.Context, message *domain.Message) error {
	return withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		sql := "INSERT INTO messages (id, sender, from_support, ticket_id, text, created_at) values ($1, $2, $3, $4, $5, $6)"
		_, err := repo.pool.Exec(ctx, sql, message.Id, message.Sender, message.FromSupport, message.TicketId, message.Text, message.CreatedAt)
		return err
	})
}

func (repo *ChatRepo) RemoveMessageById(ctx context.Context, id string) error {
	return withTx(ctx, repo.pool, func(tx pgx.Tx) error {
		sql := "DELETE FROM messages WHERE id = $1"
		_, err := repo.pool.Exec(ctx, sql, id)
		return err
	})
}
