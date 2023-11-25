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
}
