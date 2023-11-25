package repository

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(name,description,category_id,owner_id) VALUES ($1,$2,$3,$4) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, product.Name, product.Description, product.CategoryID, product.OwnerID)
	err := row.Scan(&product.ID)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name=$1, description=$2, category_id=$3 WHERE id = $4"

	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.CategoryID, product.ID)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "UPDATE products SET deletedAt=now() WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, product.ID)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, productID int) (domain.Product, error) {
	SQL := "SELECT id,name,description,category_id,owner_id FROM products WHERE id = $1 AND deletedAt IS NULL"

	product := domain.Product{}

	err := tx.QueryRowContext(ctx, SQL, productID).Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.OwnerID)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, err // product not found
		}
		return domain.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT id,name,description,category_id,owner_id FROM products WHERE deletedAt IS NULL"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.OwnerID)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
