package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GalichAnton/chat-server/internal/repository"
	repoMocks "github.com/GalichAnton/chat-server/internal/repository/mocks"
	chatService "github.com/GalichAnton/chat-server/internal/services/chat"
	"github.com/GalichAnton/platform_common/pkg/db"
	txMocks "github.com/GalichAnton/platform_common/pkg/db/mocks"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoryMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txTransactorMockFunc func(mc *minimock.Controller) db.Transactor

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name               string
		args               args
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				return repoMocks.NewLogRepositoryMock(mc)
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return repoMocks.NewUserRepositoryMock(mc)
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				return txMocks.NewTransactorMock(mc)
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return repoMocks.NewUserRepositoryMock(mc)
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				return repoMocks.NewLogRepositoryMock(mc)
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				return txMocks.NewTransactorMock(mc)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				chatRepositoryMock := tt.chatRepositoryMock(mc)
				userRepositoryMock := tt.userRepositoryMock(mc)
				logRepositoryMock := tt.logRepositoryMock(mc)
				txManagerMock := transaction.NewTransactionManager(tt.txTransactorMock(mc))
				service := chatService.NewService(
					chatRepositoryMock, userRepositoryMock, logRepositoryMock, txManagerMock,
				)

				err := service.Delete(tt.args.ctx, tt.args.id)
				if err != nil {
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
			},
		)
	}
}
