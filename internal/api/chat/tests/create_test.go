package tests

import (
	"context"
	"fmt"
	"testing"

	chatApi "github.com/GalichAnton/chat-server/internal/api/chat"
	"github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/services"
	serviceMocks "github.com/GalichAnton/chat-server/internal/services/mocks"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) services.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) services.MessageService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		users = []string{gofakeit.Name(), gofakeit.Name()}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Info: &desc.ChatInfo{
				Usernames: users,
			},
		}

		chatInfo = &chat.Info{
			Users: users,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *desc.CreateResponse
		err                error
		chatServiceMock    chatServiceMockFunc
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) services.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatInfo).Return(id, nil)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) services.MessageService {
				return serviceMocks.NewMessageServiceMock(mc)
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) services.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatInfo).Return(0, serviceErr)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) services.MessageService {
				return serviceMocks.NewMessageServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				chatServiceMock := tt.chatServiceMock(mc)
				messageServiceMock := tt.messageServiceMock(mc)
				api := chatApi.NewImplementation(chatServiceMock, messageServiceMock)

				response, err := api.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, response)
			},
		)
	}
}
