package chat

import (
	"context"
	"log"

	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/platform_common/pkg/db"
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
func (r *Repository) Create(ctx context.Context, chat *modelService.Info) (int64, error) {
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
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// Delete - .
func (r *Repository) Delete(ctx context.Context, id int64) error {
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

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetChats ...
func (r *Repository) GetChats(ctx context.Context) ([]int64, error) {
	builderSelect := sq.Select(colID).
		From(chatTableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.GetChats",
		QueryRaw: query,
	}

	var ids []int64
	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
