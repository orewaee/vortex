package domain

import "errors"

var (
	ErrInvalidToken        = errors.New("invalid token")
	ErrMissingClaims       = errors.New("missing claims")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrNoTickets           = errors.New("no tickets found")
	ErrTicketNotFound      = errors.New("ticket not found")
	ErrTicketAlreadyExists = errors.New("ticket already exists")
	ErrNoConnection        = errors.New("no connection found")
	ErrAccountNotFound     = errors.New("account not found")
    ErrNoMessages          = errors.New("no messages found")
)

