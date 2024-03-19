package access

import "github.com/GalichAnton/chat-server/internal/client"

type accessInterceptor struct {
	client client.AccessClient
}

// NewAccessInterceptor ...
func NewAccessInterceptor(rpcClient client.AccessClient) *accessInterceptor {
	return &accessInterceptor{
		client: rpcClient,
	}
}
