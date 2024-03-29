package message

import "time"

// Message - .
type Message struct {
	ID     int64 `ds:"id"`
	Info   Info
	SentAt time.Time `db:"sent_at"`
}

// Info - .
type Info struct {
	ChatID  int64  `db:"chat_id"`
	From    int64  `db:"user_id"`
	Content string `db:"content"`
}
