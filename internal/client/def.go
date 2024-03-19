package client

import (
	"context"

	access "github.com/GalichAnton/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AccessClient ...
type AccessClient interface {
	Check(ctx context.Context, req *access.CheckRequest) (*emptypb.Empty, error)
}
