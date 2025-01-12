package domain

import "time"

type Ticket struct {
	Id        string
	ChatId    int64
	Topic     string
	Closed    bool
	CreatedAt time.Time
}
