package rpc

type accessClient struct{}

// NewAccessClient ...
func NewAccessClient() *accessClient {
	return &accessClient{}
}
