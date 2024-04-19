package chat

import (
	"log"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Connect ...
func (i *Implementation) Connect(req *desc.ConnectRequest, stream desc.ChatV1_ConnectServer) error {
	msgChan, err := i.chatService.Connect(req.GetChatId())

	if err != nil {
		log.Println(err)
		return err
	}

	for msg := range msgChan {
		if err := stream.Send(
			&desc.MessageInfo{
				From:   msg.Info.From,
				Text:   msg.Info.Content,
				SentAt: timestamppb.New(msg.SentAt),
			},
		); err != nil {
			return err
		}
	}

	return nil
}
