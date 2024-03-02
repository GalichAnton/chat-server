package user

import (
	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/services"
	"github.com/GalichAnton/platform_common/pkg/db"
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
