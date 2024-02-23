package converter

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/user"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/user/model"
)

// ToServiceUser ...
func ToServiceUser(user *modelRepo.User) *modelService.User {
	return &modelService.User{
		ID:     user.ID,
		ChatID: user.ChatID,
		Name:   user.Name,
	}
}
