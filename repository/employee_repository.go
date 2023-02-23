package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee
	FindById(ctx context.Context, tx *sql.Tx, employeeId int, branchName string) (domain.Employee, error)
	Create(ctx context.Context, tx *sql.Tx, employee *domain.Employee)
	Update(ctx context.Context, tx *sql.Tx, employee *domain.Employee)
	Delete(ctx context.Context, tx *sql.Tx, employeeId int)
}
