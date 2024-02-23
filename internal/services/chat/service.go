package user

import (
	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/services"
)

var _ services.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
	userRepository repository.UserRepository
}

// NewService ...
func NewService(chatRepository repository.ChatRepository, userRepository repository.UserRepository) *service {
	return &service{
		chatRepository: chatRepository,
		userRepository: userRepository,
	}
}
