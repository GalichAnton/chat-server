package user

import (
	"github.com/GalichAnton/chat-server/internal/client/db"
	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/services"
)

var _ services.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(
	chatRepository repository.ChatRepository,
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) *service {
	return &service{
		chatRepository: chatRepository,
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}
