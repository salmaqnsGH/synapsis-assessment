package domain

import "time"

type Transaction struct {
	ID         int
	Quantity   int
	Price      int
	TotalPrice int
	IsInCart   bool
	UserID     int
	ProductID  int
	OwnerID    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
