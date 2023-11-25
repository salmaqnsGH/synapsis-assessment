package domain

import "time"

type User struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
