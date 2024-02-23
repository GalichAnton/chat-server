package chat

import (
	"context"
	"log"

	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	chatTableName = "chat"
	colID         = "id"
	colOwner      = "owner"
)

// Repository - .
type Repository struct {
	pool *pgxpool.Pool
}

// NewChatRepository - .
func NewChatRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
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

	var chatID int64
	err = c.pool.QueryRow(ctx, query, args...).Scan(&chatID)
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

	_, err = c.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
