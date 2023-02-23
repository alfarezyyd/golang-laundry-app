package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepositoryImpl() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{}
}

func (employeeRepository *EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := "SELECT id, code, level, full_name, status FROM employees"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)
	var allEmployeeData []domain.Employee
	for rows.Next() {
		employeeData := domain.Employee{}
		err = rows.Scan(
			&employeeData.Id,
			&employeeData.Code,
			&employeeData.Level,
			&employeeData.FullName,
			&employeeData.Status,
		)
		helper.PanicIfError(err)
		allEmployeeData = append(allEmployeeData, employeeData)
	}
	return allEmployeeData
}

func (employeeRepository *EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, employeeId int, branchName string) (domain.Employee, error) {
	SQL := "SELECT e.id_address, e.id_branch, e.code, e.level, e.full_name, e.password, e.email, e.telephone_number, e.photo, e.status, e.created_at, e.updated_at, branchs.name FROM employees AS e JOIN branchs ON branchs.id = e.id WHERE e.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	defer rows.Close()
	helper.PanicIfError(err)
	var employeeData domain.Employee
	if rows.Next() {
		err = rows.Scan(
			&employeeData.IdAddress,
			&employeeData.IdBranch,
			&employeeData.Code,
			&employeeData.Level,
			&employeeData.FullName,
			&employeeData.Password,
			&employeeData.Email,
			&employeeData.TelephoneNumber,
			&employeeData.Photo,
			&employeeData.Status,
			&employeeData.CreatedAt,
			&employeeData.UpdatedAt,
			&branchName,
		)
		helper.PanicIfError(err)
		return employeeData, nil
	} else {
		return employeeData, errors.New("employee not found")
	}

}

func (employeeRepository *EmployeeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, employee *domain.Employee) {
	SQL := "INSERT INTO employees(id_address, id_branch, code, level, full_name, password, email, telephone_number, photo, status) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		employee.IdAddress,
		employee.IdBranch,
		employee.Code,
		employee.Level,
		employee.FullName,
		employee.Password,
		employee.Email,
		employee.TelephoneNumber,
		employee.Photo,
		employee.Status,
	)
	helper.PanicIfError(err)

	newEmployeeId, err := result.LastInsertId()
	helper.PanicIfError(err)

	employee.Id = int(newEmployeeId)
}

func (employeeRepository *EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee *domain.Employee) {
	SQL := "UPDATE employees SET id_branch = ?, code = ?, level = ?, full_name = ?, password = ?, telephone_number = ?, photo = ?, status = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		employee.IdBranch,
		employee.Code,
		employee.Level,
		employee.FullName,
		employee.Password,
		employee.TelephoneNumber,
		employee.Photo,
		employee.Status,
		employee.UpdatedAt,
		employee.Id,
	)
	helper.PanicIfError(err)
}

func (employeeRepository *EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employeeId int) {
	SQL := "DELETE FROM employees WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
}
