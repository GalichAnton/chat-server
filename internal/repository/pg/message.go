package pg

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/models/message"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	messageTableName = "message"
)

// MessageRepository - .
type MessageRepository struct {
	pool *pgxpool.Pool
}

// NewMessageRepository - .
func NewMessageRepository(pool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

// SendMessage - .
func (m *MessageRepository) SendMessage(ctx context.Context, message *message.Info) error {
	builderInsert := sq.Insert(messageTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(message.ChatID, message.From, message.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	var messageID int64
	err = m.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	if err != nil {
		return err
	}

	return nil
}
