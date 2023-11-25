package web

type ProductUpdateRequest struct {
	ID          int    `validate:"required" json:"id"`
	Name        string `validate:"required,max=100,min=1" json:"name"`
	Description string `json:"description"`
	CategoryID  int    `validate:"required" json:"category_id"`
}
