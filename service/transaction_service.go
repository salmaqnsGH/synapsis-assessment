package service

import (
	"context"
	"salmaqnsGH/sysnapsis-assessment/model/web"
)

type TransactionService interface {
	AddToCart(ctx context.Context, req web.CartCreateRequest) web.CartCreateResponse
}
