package interceptor

import (
	"context"

	authPB "github.com/GalichAnton/auth/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	accessToken = ""
	authAddress = "localhost:50051"
)

// AuthInterceptor ...
func AuthInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.DialContext(ctx, authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	authClient := authPB.NewAccessV1Client(conn)

	_, err = authClient.Check(
		ctx, &authPB.CheckRequest{
			EndpointAddress: info.FullMethod,
		},
	)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
