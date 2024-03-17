package rpc

import (
	"context"

	authPB "github.com/GalichAnton/auth/pkg/access_v1"
	"github.com/GalichAnton/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const authAddress = "localhost:50051"

func (c *accessClient) AccessClient(ctx context.Context) (authPB.AccessV1Client, error) {
	conn, err := grpc.DialContext(ctx, authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	closer.Add(conn.Close)

	authClient := authPB.NewAccessV1Client(conn)

	return authClient, nil
}
