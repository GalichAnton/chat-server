package user

import (
	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/services"
)

var _ services.MessageService = (*service)(nil)

type service struct {
	messageRepository repository.MessageRepository
}

// NewService ...
func NewService(messageRepository repository.MessageRepository) *service {
	return &service{
		messageRepository: messageRepository,
	}
}
