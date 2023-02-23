package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Create(ctx context.Context, tx *sql.Tx, user *domain.User)
	Update(ctx context.Context, tx *sql.Tx, user *domain.User)
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindAuthUser(ctx context.Context, tx *sql.Tx, userEmail, userPassword string) (domain.User, error)
}
