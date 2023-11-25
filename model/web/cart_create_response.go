package web

type CartCreateResponse struct {
	ID         int
	Quantity   int  `json:"quantity"`
	Price      int  `json:"price"`
	TotalPrice int  `json:"total_price"`
	IsInCart   bool `json:"is_in_cart"`
	UserID     int  `json:"user_id"`
	ProductID  int  `json:"product_id"`
	OwnerID    int  `json:"owner_id"`
}
