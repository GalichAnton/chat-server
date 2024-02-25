package chat

import (
	"github.com/GalichAnton/chat-server/internal/services"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService    services.ChatService
	messageService services.MessageService
}

// NewImplementation ...
func NewImplementation(chatService services.ChatService, messageService services.MessageService) *Implementation {
	return &Implementation{
		chatService:    chatService,
		messageService: messageService,
	}
}
