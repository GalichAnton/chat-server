package model

// User - .
type User struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	ChatID int64  `db:"chat_id"`
}
