package repository

import (
	"context"
	"database/sql"
	"salmaqnsGH/sysnapsis-assessment/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateBalance(ctx context.Context, tx *sql.Tx, userID int, balance int)
	FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error)
}
