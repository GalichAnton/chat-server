package root

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

func createChat(ctx context.Context, address string, users []string) (int64, error) {
	accessToken, err := readToken()
	if err != nil {
		log.Printf("failed to read token: %v", err)
		return 0, err
	}

	client, err := chatServerClient(address)
	if err != nil {
		return 0, err
	}

	claims, err := getTokenClaims(accessToken)
	if err != nil {
		return 0, err
	}

	err = isTokenExpired(claims)
	if err != nil {
		return 0, err
	}

	md := metadata.New(map[string]string{"Authorization": tokenHeader + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Create(
		ctx, &desc.CreateRequest{
			Info: &desc.ChatInfo{
				Usernames: users,
			},
		},
	)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}
