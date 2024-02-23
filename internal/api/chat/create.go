package chat

import (
	"context"
	"log"

	"github.com/GalichAnton/chat-server/internal/converter/chat"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

// Create ...
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, chat.ToServiceChatInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("created chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
