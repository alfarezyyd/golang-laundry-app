package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type BranchRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Branch
	FindById(ctx context.Context, tx *sql.Tx, branchId int) (domain.Branch, error)
	Create(ctx context.Context, tx *sql.Tx, branch *domain.Branch)
	Update(ctx context.Context, tx *sql.Tx, branch *domain.Branch)
	Delete(ctx context.Context, tx *sql.Tx, branchId int)
}
