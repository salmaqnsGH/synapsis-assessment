package repository

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type TransactionRepository interface {
	AddToCart(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	IsProductInCart(ctx context.Context, tx *sql.Tx, productID int, userID int) (domain.Transaction, error)
	UpdateByID(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	FindAllProductInCart(ctx context.Context, tx *sql.Tx, userID int) []domain.Transaction
	DeleteProductInCart(ctx context.Context, tx *sql.Tx, cartID int)
	GetProductByCartID(ctx context.Context, tx *sql.Tx, cartID int) (domain.Transaction, error)
	SetCartFalse(ctx context.Context, tx *sql.Tx, cartID int)
}
