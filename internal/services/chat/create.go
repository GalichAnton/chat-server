package user

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	chatUser "github.com/GalichAnton/chat-server/internal/models/user"
)

// Create ...
func (s *service) Create(ctx context.Context, chat *modelService.Info) (int64, error) {
	id, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		return 0, err
	}

	for _, user := range chat.Users {
		newUser := chatUser.User{
			Name:   user,
			ChatID: id,
		}
		_, err = s.userRepository.Create(ctx, &newUser)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
