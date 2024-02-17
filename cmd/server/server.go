package server

import (
	"context"

	"github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/models/message"
	chatUser "github.com/GalichAnton/chat-server/internal/models/user"
	"github.com/GalichAnton/chat-server/internal/repository"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ChatServer - .
type ChatServer struct {
	desc.UnimplementedChatV1Server
	chatRepository    repository.ChatRepository
	userRepository    repository.UserRepository
	messageRepository repository.MessageRepository
}

// NewChatServer - .
func NewChatServer(chatRepository repository.ChatRepository, userRepository repository.UserRepository,
	messageRepository repository.MessageRepository) *ChatServer {
	return &ChatServer{
		chatRepository:    chatRepository,
		userRepository:    userRepository,
		messageRepository: messageRepository,
	}
}

// Create - .
func (s *ChatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	newChat := chat.Info{
		Owner: req.GetOwner(),
		Users: req.GetUsernames(),
	}

	id, err := s.chatRepository.Create(ctx, &newChat)
	if err != nil {
		return nil, err
	}

	for _, user := range newChat.Users {
		newUser := chatUser.User{
			Name:   user,
			ChatID: id,
		}
		_, err := s.userRepository.Create(ctx, &newUser)
		if err != nil {
			return nil, err
		}
	}
	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Delete - .
func (s *ChatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatRepository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SendMessage - .
func (s *ChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	newMessage := message.Info{
		ChatID: req.GetChatId(),
		From:   req.GetFrom(),
		Text:   req.GetText(),
	}

	err := s.messageRepository.SendMessage(ctx, &newMessage)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
