package access

import "github.com/GalichAnton/chat-server/internal/client"

type accessInterceptor struct {
	client client.RPCClient
}

// NewAccessInterceptor ...
func NewAccessInterceptor(rpcClient client.RPCClient) *accessInterceptor {
	return &accessInterceptor{
		client: rpcClient,
	}
}
