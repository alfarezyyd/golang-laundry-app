package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type BranchRepositoryImpl struct {
}

func NewBranchRepositoryImpl() *BranchRepositoryImpl {
	return &BranchRepositoryImpl{}
}

func (branchRepository *BranchRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Branch {
	SQL := "SELECT id, name, telephone_number, created_at, updated_at FROM branchs"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allBranchData []domain.Branch
	for rows.Next() {
		branchData := domain.Branch{}
		err = rows.Scan(
			&branchData.Id,
			&branchData.Name,
			&branchData.TelephoneNumber,
			&branchData.CreatedAt,
			&branchData.UpdatedAt,
		)
		helper.PanicIfError(err)
		allBranchData = append(allBranchData, branchData)
	}
	return allBranchData
}

func (branchRepository *BranchRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, branchId int) (domain.Branch, error) {
	SQL := "SELECT id, id_address, name, telephone_number, created_at, updated_at FROM branchs WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, branchId)
	defer rows.Close()
	helper.PanicIfError(err)

	var branchData domain.Branch
	if rows.Next() {
		err = rows.Scan(
			&branchData.Id,
			&branchData.IdAddress,
			&branchData.Name,
			&branchData.TelephoneNumber,
			&branchData.CreatedAt,
			&branchData.UpdatedAt,
		)
		helper.PanicIfError(err)
		return branchData, nil
	} else {
		return branchData, errors.New("branch not found")
	}

}

func (branchRepository *BranchRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, branch *domain.Branch) {
	SQL := "INSERT INTO branchs(id_address, name, telephone_number) VALUES (?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		branch.IdAddress,
		branch.Name,
		branch.TelephoneNumber,
	)
	helper.PanicIfError(err)
	newBranchId, err := result.LastInsertId()
	helper.PanicIfError(err)
	branch.Id = int(newBranchId)
}

func (branchRepository *BranchRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, branch *domain.Branch) {
	SQL := "UPDATE branchs SET name = ?, telephone_number = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		branch.Name,
		branch.TelephoneNumber,
		branch.UpdatedAt,
		branch.Id,
	)
	helper.PanicIfError(err)
}

func (branchRepository *BranchRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, branchId int) {
	SQL := "DELETE FROM branchs WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, branchId)
	helper.PanicIfError(err)
}
