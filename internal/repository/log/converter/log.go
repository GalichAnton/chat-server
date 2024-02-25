package converter

import (
	modelService "github.com/GalichAnton/chat-server/internal/models/log"
	modelRepo "github.com/GalichAnton/chat-server/internal/repository/log/model"
)

// ToServiceLog ...
func ToServiceLog(log *modelRepo.Log) *modelService.Log {
	return &modelService.Log{
		ID:        log.ID,
		Info:      ToServiceLogInfo(log.Info),
		CreatedAt: log.CreatedAt,
	}
}

// ToServiceLogInfo ...
func ToServiceLogInfo(info modelRepo.LogInfo) modelService.Info {
	return modelService.Info{
		Action:     info.Action,
		EntityID:   info.EntityID,
		EntityType: info.EntityType,
	}
}
