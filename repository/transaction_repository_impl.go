package repository

import (
	"context"
	"database/sql"
	"errors"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `
	INSERT INTO transactions(
		quantity,
		price,
		total_price,
		user_id,
		product_id,
		owner_id
	) 
	VALUES ($1,$2,$3,$4,$5,$6) 
	RETURNING id,is_in_cart`

	row := tx.QueryRowContext(
		ctx,
		SQL,
		transaction.Quantity,
		transaction.Price,
		transaction.TotalPrice,
		transaction.UserID,
		transaction.ProductID,
		transaction.OwnerID,
	)
	err := row.Scan(&transaction.ID, &transaction.IsInCart)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) IsProductInCart(ctx context.Context, tx *sql.Tx, productID int, userID int) (domain.Transaction, error) {
	SQL := `
	SELECT 
		id,
		quantity,
		price,
		total_price,
		user_id,
		product_id,
		owner_id,
		is_in_cart
	FROM transactions 
	WHERE product_id = $1 AND user_id = $2 AND is_in_cart = true AND deleted_at IS NULL`
	rows, err := tx.QueryContext(ctx, SQL, productID, userID)
	helper.PanicIfError(err)

	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(&transaction.ID, &transaction.Quantity, &transaction.Price, &transaction.TotalPrice, &transaction.UserID, &transaction.ProductID, &transaction.OwnerID, &transaction.IsInCart)
		helper.PanicIfError(err)
		return transaction, nil
	} else {
		return transaction, errors.New("transaction is not found")
	}
}

func (repository *TransactionRepositoryImpl) UpdateByID(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `UPDATE transactions 
	SET
	quantity = $1,
	price = $2,
	total_price = $3,
	user_id = $4,
	product_id = $5,
	owner_id = $6,
	updated_at=now() WHERE id = $7`

	_, err := tx.ExecContext(
		ctx,
		SQL,
		&transaction.Quantity,
		&transaction.Price,
		&transaction.TotalPrice,
		&transaction.UserID,
		&transaction.ProductID,
		&transaction.OwnerID,
		&transaction.ID,
	)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) FindAllProductInCart(ctx context.Context, tx *sql.Tx, userID int) []domain.Transaction {
	SQL := `
	SELECT 
		id,
		product_id,
		quantity,
		price,
		total_price,
		user_id,
		product_id,
		owner_id,
		is_in_cart
	FROM transactions 
	WHERE user_id = $1 AND is_in_cart = true AND deleted_at IS NULL`
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)

	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.ProductID, &transaction.Quantity, &transaction.Price, &transaction.TotalPrice, &transaction.UserID, &transaction.ProductID, &transaction.OwnerID, &transaction.IsInCart)
		helper.PanicIfError(err)

		transactions = append(transactions, transaction)
	}

	return transactions
}

func (repository *TransactionRepositoryImpl) DeleteProductInCart(ctx context.Context, tx *sql.Tx, cartID int) {
	SQL := "UPDATE transactions SET deleted_at=now() WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, cartID)
	helper.PanicIfError(err)
}
