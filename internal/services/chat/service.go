package chat

import (
	"sync"

	"github.com/GalichAnton/chat-server/internal/repository"
	"github.com/GalichAnton/chat-server/internal/repository/message/model"
	"github.com/GalichAnton/chat-server/internal/services"
	"github.com/GalichAnton/platform_common/pkg/db"
)

const messagesBuffer = 100

var _ services.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager

	channels   map[int64]chan *model.Message
	mxChannels sync.RWMutex

	chats  map[int64]*chat
	mxChat sync.RWMutex
}

type chat struct {
	streams map[string]*model.Message
	m       sync.RWMutex
}

// NewService ...
func NewService(
	chatRepository repository.ChatRepository,
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) *service {
	return &service{
		chatRepository: chatRepository,
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
		chats:          make(map[int64]*chat),
		channels:       make(map[int64]chan *model.Message),
	}
}
