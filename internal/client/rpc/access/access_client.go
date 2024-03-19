package access

import (
	"context"

	access "github.com/GalichAnton/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type accessClient struct {
	client access.AccessV1Client
}

// NewAccessClient ...
func NewAccessClient(client access.AccessV1Client) *accessClient {
	return &accessClient{
		client: client,
	}
}

func (a *accessClient) Check(ctx context.Context, req *access.CheckRequest) (*emptypb.Empty, error) {
	_, err := a.client.Check(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
