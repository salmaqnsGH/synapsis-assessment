package web

type UserUpdateRequest struct {
	ID       int    `validate:"required" json:"id"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=1" json:"password"`
	Balance  int    `json:"balance"`
}
