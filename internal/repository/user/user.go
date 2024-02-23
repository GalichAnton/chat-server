package user

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	userTableName = "chat_user"
	colName       = "name"
	colChatID     = "chat_id"
)

// Repository - .
type Repository struct {
	pool *pgxpool.Pool
}

// NewUserRepository - .
func NewUserRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
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

	var userID int64
	err = u.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
