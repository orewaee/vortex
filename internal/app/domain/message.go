package domain

import "time"

type Message struct {
	Id          string
	Sender      string
	FromSupport bool
	TicketId    string
	Text        string
	CreatedAt   time.Time
}
