package message

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/message"
	"github.com/GalichAnton/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	messageTableName = "message"
	colChatID        = "chat_id"
	colUserID        = "user_id"
	colContent       = "content"
)

// Repository - .
type Repository struct {
	db db.Client
}

// NewMessageRepository - .
func NewMessageRepository(db db.Client) *Repository {
	return &Repository{db: db}
}

// SendMessage - .
func (m *Repository) SendMessage(ctx context.Context, message *modelService.Info) error {
	builderInsert := sq.Insert(messageTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colChatID, colUserID, colContent).
		Values(message.ChatID, message.From, message.Content).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "message_repository.SendMessage",
		QueryRaw: query,
	}

	var messageID int64
	err = m.db.DB().QueryRowContext(ctx, q, args...).Scan(&messageID)
	if err != nil {
		return err
	}

	return nil
}
