package repository

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/helper"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "INSERT INTO users(name,username,password,balance) VALUES ($1,$2,$3,0) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, user.Name, user.Username, user.Password)
	err := row.Scan(&user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	const SQL = `SELECT id, name, username, password, balance FROM users WHERE username = $1 AND deleted_at IS NULL`

	user := domain.User{}
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil // user not found
		}
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name=$1, password=$2, username=$3, balance=$4, updated_at=now() WHERE id = $5"

	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Password, user.Username, user.Balance, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) UpdateBalance(ctx context.Context, tx *sql.Tx, userID int, balance int) {
	SQL := "UPDATE users SET balance=$1, updated_at=now() WHERE id = $2"

	_, err := tx.ExecContext(ctx, SQL, balance, userID)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error) {
	const SQL = `SELECT id, name, username, password, balance FROM users WHERE id = $1 AND deleted_at IS NULL`

	user := domain.User{}
	err := tx.QueryRowContext(ctx, SQL, userID).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil // user not found
		}
		return domain.User{}, err
	}

	return user, nil
}
