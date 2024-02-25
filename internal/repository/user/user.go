package user

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/client/db"
	modelService "github.com/GalichAnton/chat-server/internal/models/user"
	sq "github.com/Masterminds/squirrel"
)

const (
	userTableName = "chat_user"
	colName       = "name"
	colChatID     = "chat_id"
)

// Repository - .
type Repository struct {
	db db.Client
}

// NewUserRepository - .
func NewUserRepository(db db.Client) *Repository {
	return &Repository{db: db}
}

// Create - .
func (u *Repository) Create(ctx context.Context, user *modelService.User) (int64, error) {
	builderInsert := sq.Insert(userTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colName, colChatID).
		Values(user.Name, user.ChatID).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}
	var userID int64
	err = u.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
