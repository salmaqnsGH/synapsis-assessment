package service

import (
	"context"
	"salmaqnsGH/sysnapsis-assessment/model/web"
)

type UserService interface {
	Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, req web.UserLoginRequest) (web.UserLoginResponse, error)
	Update(ctx context.Context, req web.UserUpdateRequest) web.UserResponse
}
