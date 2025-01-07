package dtos

import "time"

type Ticket struct {
	Id        string    `json:"id"`
	ChatId    int64     `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}
