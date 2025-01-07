package domain

import "time"

type Account struct {
	Id        string
	Name      string
	Password  string
	Perms     int
	CreatedAt time.Time
}
