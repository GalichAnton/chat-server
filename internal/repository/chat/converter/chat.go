package converter

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/chat"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/chat/model"
)

// ToServiceChat ...
func ToServiceChat(chat *modelRepo.Chat) *modelService.Chat {
	return &modelService.Chat{
		ID:   chat.ID,
		Info: ToServiceChatInfo(chat.Info),
	}
}

// ToServiceChatInfo ...
func ToServiceChatInfo(info modelRepo.Info) modelService.Info {
	return modelService.Info{
		Owner: info.Owner,
		Users: info.Users,
	}
}
