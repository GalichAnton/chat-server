package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/GalichAnton/chat-server/internal/models/chat"
	"github.com/GalichAnton/chat-server/internal/models/log"
	chatUser "github.com/GalichAnton/chat-server/internal/models/user"
	"github.com/GalichAnton/chat-server/internal/repository"
	repoMocks "github.com/GalichAnton/chat-server/internal/repository/mocks"
	chatService "github.com/GalichAnton/chat-server/internal/services/chat"
	"github.com/GalichAnton/platform_common/pkg/db"
	txMocks "github.com/GalichAnton/platform_common/pkg/db/mocks"
	"github.com/GalichAnton/platform_common/pkg/db/pg"
	"github.com/GalichAnton/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoryMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txTransactorMockFunc func(mc *minimock.Controller) db.Transactor

	type args struct {
		ctx context.Context
		req *chat.Info
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    txMocks.TxMock
		id     = gofakeit.Int64()
		owner  = gofakeit.Int64()

		users = []string{gofakeit.Name(), gofakeit.Name()}

		repoErr = fmt.Errorf("repo error")
		txError = errors.Wrap(repoErr, "failed executing code inside transaction")

		req = &chat.Info{
			Owner: owner,
			Users: users,
		}

		chatInfo = &chat.Info{
			Owner: owner,
			Users: users,
		}

		logInfo = &log.Info{
			Action:     "create",
			EntityID:   id,
			EntityType: "chat",
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
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
				req: req,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), chatInfo).Return(id, nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repoMocks.NewLogRepositoryMock(mc)
				mock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), logInfo).Return(nil)
				return mock
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				for i, user := range chatInfo.Users {
					u := user
					ui := i
					newUser := &chatUser.User{
						Name:   u,
						ChatID: id,
					}
					mock.CreateMock.When(ctx, newUser).Then(int64(ui), nil)
				}

				return mock
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  txError,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), chatInfo).Return(0, repoErr)
				return mock
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return repoMocks.NewUserRepositoryMock(mc)
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				return repoMocks.NewLogRepositoryMock(mc)
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
				return mock
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
