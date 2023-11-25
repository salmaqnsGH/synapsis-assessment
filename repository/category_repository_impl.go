package repository

import (
	"context"
	"database/sql"
	"errors"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name,description) VALUES ($1,$2) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, category.Name, category.Description)
	err := row.Scan(&category.ID)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = $1, description=$2, updatedAt=now() WHERE id = $3"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Description, category.ID)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "UPDATE categories SET deletedAt=now() WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, category.ID)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	SQL := "SELECT id,name,description FROM categories WHERE id = $1 AND deletedAt IS NULL"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id,name,description FROM categories WHERE deletedAt IS NULL"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
