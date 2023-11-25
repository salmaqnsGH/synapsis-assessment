package repository

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "INSERT INTO users(name,username,password) VALUES ($1,$2,$3) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, user.Name, user.Username, user.Password)
	err := row.Scan(&user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	const SQL = `SELECT id, name, username, password FROM users WHERE username = $1`

	user := domain.User{}
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil // user not found
		}
		return domain.User{}, err
	}

	return user, nil
}
