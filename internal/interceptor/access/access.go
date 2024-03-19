package access

import (
	"context"

	authPB "github.com/GalichAnton/auth/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	accessToken = ""
)

// Access ...
func (i *accessInterceptor) Access(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err = i.client.Check(
		ctx, &authPB.CheckRequest{
			EndpointAddress: info.FullMethod,
		},
	)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
