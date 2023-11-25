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

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, req web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		OwnerID:     req.OwnerID,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		OwnerID:     product.OwnerID,
	}
}

func (service *ProductServiceImpl) Update(ctx context.Context, req web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = req.Name
	product.Description = req.Description
	product.CategoryID = req.CategoryID

	product = service.ProductRepository.Update(ctx, tx, product)

	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		OwnerID:     product.OwnerID,
	}
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productID int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, productID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindByID(ctx context.Context, productID int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, productID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		OwnerID:     product.OwnerID,
	}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponse := web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			OwnerID:     product.OwnerID,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}
