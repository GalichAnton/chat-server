package pg

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/models/message"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	messageTableName = "message"
	colChatID        = "chat_id"
	colUserID        = "user_id"
	colContent       = "content"
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
		Columns(colChatID, colUserID, colContent).
		Values(message.ChatID, message.From, message.Content).
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
