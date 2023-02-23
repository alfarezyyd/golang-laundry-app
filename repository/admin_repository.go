package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type AdminRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Admin
	FindById(ctx context.Context, tx *sql.Tx, userId int, branchName string) (domain.Admin, error)
	Create(ctx context.Context, tx *sql.Tx, user *domain.Admin)
	Update(ctx context.Context, tx *sql.Tx, user *domain.Admin)
	Delete(ctx context.Context, tx *sql.Tx, userId int)
}
