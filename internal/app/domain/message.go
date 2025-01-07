package domain

type Message struct {
	Sender      string
	FromSupport bool
	TicketId    string
	Text        string
}
