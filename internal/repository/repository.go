package repository

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/models/log"
	"github.com/GalichAnton/chat-server/internal/models/message"
	"github.com/GalichAnton/chat-server/internal/models/user"
)

// ChatRepository - .
type ChatRepository interface {
	Create(ctx context.Context, chat *chat.Info) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetChats(ctx context.Context) ([]int64, error)
}

// UserRepository - .
type UserRepository interface {
	Create(ctx context.Context, user *user.User) (int64, error)
}

// MessageRepository - .
type MessageRepository interface {
	SendMessage(ctx context.Context, message *message.Info) error
}

// LogRepository - .
type LogRepository interface {
	Create(ctx context.Context, log *log.Info) error
}
