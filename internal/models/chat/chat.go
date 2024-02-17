package chat

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
