package chat

import (
	"context"
	"errors"

	"github.com/GalichAnton/chat-server/internal/repository/message/model"
)

// InitChannels ...
func (s *service) InitChannels(ctx context.Context) error {
	ids, err := s.chatRepository.GetChats(ctx)
	if err != nil {
		return errors.New("failed to init existing chats")
	}

	for _, id := range ids {
		s.channels[id] = make(chan *model.Message, messagesBuffer)
	}

	return nil
}
