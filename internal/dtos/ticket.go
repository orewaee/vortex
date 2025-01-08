package dtos

import "time"

type Ticket struct {
	Id        string    `json:"id"`
	ChatId    int64     `json:"chat_id"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
