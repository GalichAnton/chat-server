package chat

import (
	"context"
	"log"

	"github.com/GalichAnton/chat-server/internal/client/db"
	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	sq "github.com/Masterminds/squirrel"
)

const (
	chatTableName = "chat"
	colID         = "id"
	colOwner      = "owner"
)

// Repository - .
type Repository struct {
	db db.Client
}

// NewChatRepository - .
func NewChatRepository(db db.Client) *Repository {
	return &Repository{db: db}
}

// Create - .
func (c *Repository) Create(ctx context.Context, chat *modelService.Info) (int64, error) {
	builderInsert := sq.Insert(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colOwner).
		Values(chat.Owner).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}
	var chatID int64
	err = c.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// Delete - .
func (c *Repository) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colID: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = c.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
