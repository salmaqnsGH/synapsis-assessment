package web

type UserLoginRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=1" json:"password"`
}
