package domain

import "time"

type Ticket struct {
	Id        string
	ChatId    int64
	Topic     string
	CreatedAt time.Time
}
