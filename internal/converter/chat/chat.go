package chat

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

// ToServiceChatInfo ...
func ToServiceChatInfo(info *desc.ChatInfo) *modelService.Info {
	return &modelService.Info{
		Users: info.Usernames,
	}
}

// ToStreamFromDesc ...
func ToStreamFromDesc(stream desc.ChatV1_ConnectServer) modelService.Stream {
	return stream.(modelService.Stream)
}
