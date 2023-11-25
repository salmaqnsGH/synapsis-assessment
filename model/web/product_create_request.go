package web

type ProductCreateRequest struct {
	Name        string `validate:"required,max=100,min=1" json:"name"`
	Description string `json:"description"`
	CategoryID  int    `validate:"required" json:"category_id"`
	OwnerID     int    `validate:"required" json:"owner_id"`
	Quantity    int    `validate:"required" json:"quantity"`
	Price       int    `validate:"required" json:"price"`
}
