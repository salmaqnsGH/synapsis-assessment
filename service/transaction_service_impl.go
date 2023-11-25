package service

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/exception"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
	"salmaqnsGH/sysnapsis-assessment/model/web"
	"salmaqnsGH/sysnapsis-assessment/repository"

	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	ProductRepository     repository.ProductRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) AddToCart(ctx context.Context, req web.CartCreateRequest) web.CartCreateResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, req.ProductID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cart, _ := service.TransactionRepository.IsProductInCart(ctx, tx, req.ProductID, req.UserID)

	quantity := cart.Quantity + req.Quantity
	totalPrice := product.Price * req.Quantity

	transaction := domain.Transaction{
		Quantity:   quantity,
		UserID:     req.UserID,
		ProductID:  req.ProductID,
		OwnerID:    product.OwnerID,
		Price:      product.Price,
		TotalPrice: totalPrice,
	}

	var empty = domain.Transaction{}

	if cart == empty {
		transaction = service.TransactionRepository.AddToCart(ctx, tx, transaction)
	} else {
		transactionRequest := domain.Transaction{
			ID:         cart.ID,
			Quantity:   quantity,
			UserID:     req.UserID,
			ProductID:  req.ProductID,
			OwnerID:    product.OwnerID,
			Price:      product.Price,
			TotalPrice: totalPrice,
		}
		transaction = service.TransactionRepository.UpdateByID(ctx, tx, transactionRequest)
	}

	return web.CartCreateResponse{
		ID:         transaction.ID,
		Quantity:   transaction.Quantity,
		Price:      transaction.Price,
		TotalPrice: transaction.TotalPrice,
		IsInCart:   transaction.IsInCart,
		UserID:     transaction.UserID,
		ProductID:  transaction.ProductID,
		OwnerID:    transaction.OwnerID,
	}
}

func (service *TransactionServiceImpl) FindAllProductInCart(ctx context.Context, userID int) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	products := service.TransactionRepository.FindAllProductInCart(ctx, tx, userID)

	var productResponses []web.TransactionResponse
	for _, product := range products {
		productResponse := web.TransactionResponse{
			ID:         product.ID,
			OwnerID:    product.OwnerID,
			Quantity:   product.Quantity,
			Price:      product.Price,
			TotalPrice: product.TotalPrice,
			IsInCart:   product.IsInCart,
			UserID:     product.UserID,
			ProductID:  product.ProductID,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, productID int, userID int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	cart, err := service.TransactionRepository.IsProductInCart(ctx, tx, productID, userID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TransactionRepository.DeleteProductInCart(ctx, tx, cart.ID)
}
