package converter

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/message"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/message/model"
)

// ToServiceMessage ...
func ToServiceMessage(message *modelRepo.Message) *modelService.Message {
	return &modelService.Message{
		ID:   message.ID,
		Info: ToServiceMessageInfo(message.Info),
	}
}

// ToServiceMessageInfo ...
func ToServiceMessageInfo(info modelRepo.Info) modelService.Info {
	return modelService.Info{
		ChatID:  info.ChatID,
		Content: info.Content,
	}
}
