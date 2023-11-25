package service

import (
	"context"
	"salmaqnsGH/sysnapsis-assessment/model/web"
)

type ProductService interface {
	Create(ctx context.Context, req web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, req web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productID int)
	FindByID(ctx context.Context, productID int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
