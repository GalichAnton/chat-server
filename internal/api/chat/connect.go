package chat

import (
	"log"

	converter "github.com/GalichAnton/chat-server/internal/converter/chat"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

// Connect ...
func (i *Implementation) Connect(req *desc.ConnectRequest, stream desc.ChatV1_ConnectServer) error {
	err := i.chatService.Connect(req.GetChatId(), req.GetEmail(), converter.ToStreamFromDesc(stream))
	log.Println(err)
	return err
}
