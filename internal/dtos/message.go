package dtos

import "time"

type Message struct {
	Id          string    `json:"id"`
	Sender      string    `json:"sender"`
	FromSupport bool      `json:"from_support"`
	TicketId    string    `json:"ticket_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}
