package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GalichAnton/chat-server/internal/models/message"
	"github.com/GalichAnton/chat-server/internal/repository"
	repoMocks "github.com/GalichAnton/chat-server/internal/repository/mocks"
	messageService "github.com/GalichAnton/chat-server/internal/services/message"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.MessageRepository

	type args struct {
		ctx context.Context
		req *message.Info
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId = gofakeit.Int64()
		from   = gofakeit.Int64()
		text   = gofakeit.Animal()

		repoErr = fmt.Errorf("repo error")

		messageInfo = &message.Info{
			ChatID:  chatId,
			From:    from,
			Content: text,
		}
	)

	tests := []struct {
		name                  string
		args                  args
		err                   error
		messageRepositoryMock messageRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: messageInfo,
			},
			err: nil,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, messageInfo).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: messageInfo,
			},
			err: repoErr,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, messageInfo).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				messageRepositoryMock := tt.messageRepositoryMock(mc)
				service := messageService.NewService(messageRepositoryMock)

				err := service.SendMessage(tt.args.ctx, tt.args.req)
				if err != nil {
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
			},
		)
	}
}
