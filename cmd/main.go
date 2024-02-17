package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/GalichAnton/chat-server/cmd/server"
	"github.com/GalichAnton/chat-server/internal/config"
	"github.com/GalichAnton/chat-server/internal/config/env"
	"github.com/GalichAnton/chat-server/internal/repository/pg"
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

	chatRepository := pg.NewChatRepository(pool)
	messageRepository := pg.NewMessageRepository(pool)
	userRepository := pg.NewUserRepository(pool)

	s := grpc.NewServer()
	reflection.Register(s)

	userServer := server.NewChatServer(chatRepository, userRepository, messageRepository)

	desc.RegisterChatV1Server(s, userServer)

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
