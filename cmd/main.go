package main

import (
	"context"
	"flag"
	"log"
	"net"

	chatApi "github.com/GalichAnton/chat-server/internal/api/chat"
	"github.com/GalichAnton/chat-server/internal/config"
	"github.com/GalichAnton/chat-server/internal/config/env"
	chatRepo "github.com/GalichAnton/chat-server/internal/repository/chat"
	"github.com/GalichAnton/chat-server/internal/repository/message"
	"github.com/GalichAnton/chat-server/internal/repository/user"
	chatService "github.com/GalichAnton/chat-server/internal/services/chat"
	messageService "github.com/GalichAnton/chat-server/internal/services/message"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to parse gRPC config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to parse PG config: %v", err)
	}

	ctx := context.Background()
	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	chatRepository := chatRepo.NewChatRepository(pool)
	messageRepository := message.NewMessageRepository(pool)
	userRepository := user.NewUserRepository(pool)

	chatSrv := chatService.NewService(chatRepository, userRepository)
	messageSrv := messageService.NewService(messageRepository)

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterChatV1Server(s, chatApi.NewImplementation(chatSrv, messageSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
