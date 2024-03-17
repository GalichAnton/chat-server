package client

import (
	"context"

	authPB "github.com/GalichAnton/auth/pkg/access_v1"
)

// RPCClient ...
type RPCClient interface {
	AccessClient(ctx context.Context) (authPB.AccessV1Client, error)
}
