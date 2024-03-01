package tests

import (
	"context"
	"fmt"
	"testing"

	chatUser "github.com/GalichAnton/chat-server/internal/models/user"
	"github.com/GalichAnton/chat-server/internal/repository"
	repoMocks "github.com/GalichAnton/chat-server/internal/repository/mocks"
	userService "github.com/GalichAnton/chat-server/internal/services/user"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *chatUser.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = gofakeit.Int64()
		chatId = gofakeit.Int64()
		name   = gofakeit.Name()

		repoErr = fmt.Errorf("repo error")

		user = &chatUser.User{
			ChatID: chatId,
			Name:   name,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: user,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: user,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				userRepositoryMock := tt.userRepositoryMock(mc)
				service := userService.NewService(userRepositoryMock)

				response, err := service.Create(tt.args.ctx, tt.args.req)
				if err != nil {
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.Equal(t, tt.want, response)
			},
		)
	}
}
