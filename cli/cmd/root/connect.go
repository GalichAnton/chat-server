package root

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
)

func connectChat(ctx context.Context, address string, chatID int64) error {
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
	stream, err := client.Connect(
		ctx, &desc.ConnectRequest{
			ChatId: chatID,
			Email:  claims.Email,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(color.CyanString("[Connected]"))

	go func() {
		for {
			message, errRecv := stream.Recv()
			if errRecv == io.EOF {
				return
			}
			if errRecv != nil {
				log.Printf("failed to receive message: %v", errRecv)
				return
			}

			fmt.Printf(
				"[%v] - [from: %s]: %s",
				color.YellowString(message.GetSentAt().AsTime().Format(time.RFC3339)),
				color.BlueString(message.GetFrom()),
				message.GetText(),
			)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		fmt.Print(lineUp)
		_, err = client.SendMessage(
			ctx, &desc.SendMessageRequest{
				ChatId: chatID,
				Info: &desc.MessageInfo{
					From:   claims.Email,
					Text:   text,
					SentAt: timestamppb.Now(),
				},
			},
		)
		if err != nil {
			log.Printf("failed to send message: %v", err)
		}
	}
}
