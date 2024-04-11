package root

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	descAuth "github.com/GalichAnton/auth/pkg/auth_v1"
	descChat "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"github.com/GalichAnton/platform_common/pkg/closer"
)

func authClient(address string) (descAuth.AuthV1Client, error) {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	closer.Add(conn.Close)

	return descAuth.NewAuthV1Client(conn), nil
}

func chatServerClient(address string) (descChat.ChatV1Client, error) {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	closer.Add(conn.Close)

	return descChat.NewChatV1Client(conn), nil
}
