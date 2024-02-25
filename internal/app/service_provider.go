package app

import (
	"context"
	"log"

	"github.com/GalichAnton/chat-server/internal/api/chat"
	"github.com/GalichAnton/chat-server/internal/client/db"
	"github.com/GalichAnton/chat-server/internal/client/db/pg"
	"github.com/GalichAnton/chat-server/internal/client/db/transaction"
	"github.com/GalichAnton/chat-server/internal/closer"
	"github.com/GalichAnton/chat-server/internal/config"
	"github.com/GalichAnton/chat-server/internal/config/env"
	"github.com/GalichAnton/chat-server/internal/repository"
	chatRepository "github.com/GalichAnton/chat-server/internal/repository/chat"
	logRepository "github.com/GalichAnton/chat-server/internal/repository/log"
	messageRepository "github.com/GalichAnton/chat-server/internal/repository/message"
	userRepository "github.com/GalichAnton/chat-server/internal/repository/user"
	"github.com/GalichAnton/chat-server/internal/services"
	chatService "github.com/GalichAnton/chat-server/internal/services/chat"
	messageService "github.com/GalichAnton/chat-server/internal/services/message"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient          db.Client
	txManager         db.TxManager
	userRepository    repository.UserRepository
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
	logRepository     repository.LogRepository

	chatService    services.ChatService
	messageService services.MessageService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewChatRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewMessageRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewLogRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) services.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.UserRepository(ctx),
			s.LogRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) MessageService(ctx context.Context) services.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(s.MessageRepository(ctx))
	}

	return s.messageService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx), s.MessageService(ctx))
	}

	return s.chatImpl
}
