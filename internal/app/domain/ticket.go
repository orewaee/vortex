package domain

import "time"

type Ticket struct {
	Id        string
	ChatId    int64
	CreatedAt time.Time
}
