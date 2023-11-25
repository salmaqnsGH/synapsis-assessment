package service

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
	"salmaqnsGH/sysnapsis-assessment/model/web"
	"salmaqnsGH/sysnapsis-assessment/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	helper.PanicIfError(err)

	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(passwordHash),
	}

	user, err = service.UserRepository.Save(ctx, tx, user)
	helper.PanicIfError(err)

	return web.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) (web.UserLoginResponse, error) {
	err := service.Validate.Struct(req)
	if err != nil {
		return web.UserLoginResponse{}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, req.Username)
	if err != nil {
		return web.UserLoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return web.UserLoginResponse{}, err
	}

	token, err := helper.GenerateToken(user.ID)
	helper.PanicIfError(err)

	return web.UserLoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}, nil
}
