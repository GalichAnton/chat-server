package chat

import (
	"context"
	"log"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete ...
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	err := i.chatService.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("deleted chat with id: %d", id)

	return &emptypb.Empty{}, nil
}
