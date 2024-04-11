package converter

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/message"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/message/model"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// ToMessageFromService ...
func ToMessageFromService(message *modelRepo.Message) *desc.MessageInfo {
	return &desc.MessageInfo{
		From:   message.Info.From,
		Text:   message.Info.Content,
		SentAt: timestamppb.New(message.SentAt),
	}
}
