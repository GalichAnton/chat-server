package message

import "time"

// Message - .
type Message struct {
	ID        int64 `ds:"id"`
	Info      Info
	Timestamp time.Time `db:"timestamp"`
}

// Info - .
type Info struct {
	ChatID int64  `db:"chat_id"`
	From   int64  `db:"user_id"`
	Text   string `db:"text"`
}
