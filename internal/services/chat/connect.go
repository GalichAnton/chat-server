package chat

import (
	"errors"

	modelService "github.com/GalichAnton/chat-server/internal/models/message"
	"github.com/GalichAnton/chat-server/internal/repository/message/converter"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/message/model"
)

func (s *service) Connect(chatID int64) (
	<-chan *modelService.Message,
	error,
) {
	s.mxChannels.RLock()
	chatChan, ok := s.channels[chatID]
	s.mxChannels.RUnlock()

	if !ok {
		return nil, errors.New("chat not found")
	}

	s.mxChat.Lock()
	if _, okChat := s.chats[chatID]; !okChat {
		s.chats[chatID] = &chat{
			streams: make(map[string]*modelRepo.Message),
		}
	}

	s.mxChat.Unlock()

	msgChan := make(chan *modelService.Message)

	go func() {
		for {
			select {
			case msg, okCh := <-chatChan:

				if !okCh {
					close(msgChan)
					return
				}

				info := converter.ToServiceMessage(msg)

				msgChan <- info
			}
		}
	}()

	return msgChan, nil
}
