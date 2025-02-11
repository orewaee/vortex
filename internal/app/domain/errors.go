package domain

import "errors"

var (
	ErrNoTicket = errors.New("no ticket found")

	// todo remove 'already' from name

	ErrTicketAlreadyExists = errors.New("ticket already exists")
	ErrTicketAlreadyOpen   = errors.New("ticket already open")
	ErrTicketAlreadyClosed = errors.New("ticket already closed")

	ErrNoTickets      = errors.New("no tickets found")
	ErrTicketNotFound = errors.New("ticket not found")

	ErrInvalidToken       = errors.New("invalid token")
	ErrMissingClaims      = errors.New("missing claims")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNoConnection       = errors.New("no connection found")
	ErrAccountNotFound    = errors.New("account not found")
	ErrNoMessages         = errors.New("no messages found")
)
