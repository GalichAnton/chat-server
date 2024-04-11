package chat

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/models/log"
	chatUser "github.com/GalichAnton/chat-server/internal/models/user"
)

// Create ...
func (s *service) Create(ctx context.Context, chat *modelService.Info) (int64, error) {
	var newChatID int64

	err := s.txManager.ReadCommitted(
		ctx, func(ctx context.Context) error {
			id, errTx := s.chatRepository.Create(ctx, chat)
			if errTx != nil {
				return errTx
			}

			newChatID = id
			newLog := log.Info{
				Action:     "create",
				EntityID:   id,
				EntityType: "chat",
			}

			errTx = s.logRepository.Create(ctx, &newLog)
			if errTx != nil {
				return errTx
			}

			return nil
		},
	)
	if err != nil {
		return 0, err
	}

	for _, user := range chat.Users {
		newUser := chatUser.User{
			Name:   user,
			ChatID: newChatID,
		}
		_, err = s.userRepository.Create(ctx, &newUser)
		if err != nil {
			return 0, err
		}
	}

	return newChatID, nil
}
