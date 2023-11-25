package domain

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	CategoryID  int
	OwnerID     int
	Quantity    int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
