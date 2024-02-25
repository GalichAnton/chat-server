package user

import (
	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/services"
)

var _ services.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

// NewService ...
func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}
