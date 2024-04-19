package root

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

func deleteChat(ctx context.Context, address string, chatID int64) error {
	accessToken, err := readToken()
	if err != nil {
		log.Printf("failed to read token: %v", err)
		return err
	}

	client, err := chatServerClient(address)
	if err != nil {
		return err
	}

	claims, err := getTokenClaims(accessToken)
	if err != nil {
		return err
	}

	err = isTokenExpired(claims)
	if err != nil {
		return err
	}

	md := metadata.New(map[string]string{"Authorization": tokenHeader + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = client.Delete(
		ctx, &desc.DeleteRequest{
			Id: chatID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
