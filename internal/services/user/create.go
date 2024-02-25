package user

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/user"
)

// Create ...
func (s *service) Create(ctx context.Context, user *modelService.User) (int64, error) {
	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
