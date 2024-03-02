package tests

import (
	"context"
	"fmt"
	"testing"

	chatApi "github.com/GalichAnton/chat-server/internal/api/chat"
	"github.com/GalichAnton/chat-server/internal/services"
	serviceMocks "github.com/GalichAnton/chat-server/internal/services/mocks"
	desc "github.com/GalichAnton/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) services.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) services.MessageService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
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
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
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

				response, err := api.Delete(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, response)
			},
		)
	}
}
