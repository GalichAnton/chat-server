package chat

import (
	"errors"

	chatModel "github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/repository/message/converter"
)

func (s *service) Connect(chatID int64, username string, stream chatModel.Stream) error {
	s.mxChannels.RLock()
	chatChan, ok := s.channels[chatID]
	s.mxChannels.RUnlock()

	if !ok {
		return errors.New("chat not found")
	}

	s.mxChat.Lock()
	if _, okChat := s.chats[chatID]; !okChat {
		s.chats[chatID] = &chat{
			streams: make(map[string]chatModel.Stream),
		}
	}
	s.mxChat.Unlock()

	// Set stream for user
	s.chats[chatID].m.Lock()
	s.chats[chatID].streams[username] = stream
	s.chats[chatID].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			// Check if channel is closed
			if !okCh {
				return nil
			}

			// Send message for everyone in chat
			for _, st := range s.chats[chatID].streams {
				if err := st.Send(converter.ToMessageFromService(msg)); err != nil {
					return err
				}
			}
		case <-stream.Context().Done():
			// Delete stream for user when context is dead
			s.chats[chatID].m.Lock()
			delete(s.chats[chatID].streams, username)
			s.chats[chatID].m.Unlock()
			return nil
		}
	}
}
