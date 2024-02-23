package services

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/models/message"
	"github.com/GalichAnton/chat-server/internal/models/user"
)

// UserService ...
type UserService interface {
	Create(ctx context.Context, user *user.User) (int64, error)
}

// ChatService ...
type ChatService interface {
	Create(ctx context.Context, chat *chat.Info) (int64, error)
	Delete(ctx context.Context, id int64) error
}

// MessageService ...
type MessageService interface {
	SendMessage(ctx context.Context, message *message.Info) error
}
