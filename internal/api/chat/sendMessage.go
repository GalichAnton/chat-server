package chat

import (
	"context"
	"log"

	"github.com/GalichAnton/chat-server/internal/converter/message"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage ...
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.messageService.SendMessage(ctx, message.ToServiceMessageInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("created message")

	return &emptypb.Empty{}, nil
}
