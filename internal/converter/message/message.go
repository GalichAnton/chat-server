package message

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/message"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

// ToServiceMessageInfo ...
func ToServiceMessageInfo(info *desc.MessageInfo) *modelService.Info {
	return &modelService.Info{
		From:    info.From,
		Content: info.Text,
	}
}
