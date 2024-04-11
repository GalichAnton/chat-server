package chat

import desc "github.com/GalichAnton/chat-server/pkg/chat_v1"

// Chat - .
type Chat struct {
	ID   int64
	Info Info
}

// Info - .
type Info struct {
	Owner int64
	Users []string
}

// Stream ...
type Stream interface {
	desc.ChatV1_ConnectServer
}
