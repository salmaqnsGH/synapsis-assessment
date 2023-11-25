package web

type UserCreateRequest struct {
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=1" json:"password"`
}
