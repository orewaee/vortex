package dtos

import "time"

type Ticket struct {
	Id        string    `json:"id"`
	ChatId    int64     `json:"chat_id"`
	Topic     string    `json:"topic"`
	Closed    bool      `json:"closed"`
	CreatedAt time.Time `json:"created_at"`
}
