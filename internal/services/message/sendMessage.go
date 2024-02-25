package user

import (
	"context"

	modelService "github.com/GalichAnton/chat-server/internal/models/message"
)

// SendMessage ...
func (s *service) SendMessage(ctx context.Context, message *modelService.Info) error {
	err := s.messageRepository.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
