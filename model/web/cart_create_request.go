package web

type CartCreateRequest struct {
	Quantity   int `validate:"required" json:"quantity"`
	Price      int `json:"price"`
	TotalPrice int `json:"total_price"`
	UserID     int `validate:"required" json:"user_id"`
	ProductID  int `validate:"required" json:"product_id"`
	OwnerID    int `json:"owner_id"`
}
