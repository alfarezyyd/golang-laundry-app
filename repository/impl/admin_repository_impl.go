package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepositoryImpl() *AdminRepositoryImpl {
	return &AdminRepositoryImpl{}
}

func (adminRepository *AdminRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Admin {
	SQL := "SELECT id, code, level, full_name, status FROM admins"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allAdminData []domain.Admin
	for rows.Next() {
		adminData := domain.Admin{}
		err = rows.Scan(
			&adminData.Id,
			&adminData.Code,
			&adminData.Level,
			&adminData.FullName,
			&adminData.Status,
		)
		helper.PanicIfError(err)
		allAdminData = append(allAdminData, adminData)
	}
	return allAdminData
}

func (adminRepository *AdminRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, adminId int, branchName string) (domain.Admin, error) {
	SQL := "SELECT a.id, a.id_address, a.id_branch, a.code, a.level, a.full_name, a.password, a.email, a.telephone_number, a.photo, a.status, a.created_at, a.updated_at, branchs.name FROM admins AS a JOIN branchs ON a.id_branch = branchs.id WHERE a.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, adminId)
	defer rows.Close()
	helper.PanicIfError(err)

	var adminData domain.Admin
	if rows.Next() {
		err = rows.Scan(
			&adminData.Id,
			&adminData.IdAddress,
			&adminData.IdBranch,
			&adminData.Code,
			&adminData.Level,
			&adminData.FullName,
			&adminData.Password,
			&adminData.Email,
			&adminData.TelephoneNumber,
			&adminData.Photo,
			&adminData.Status,
			&adminData.CreatedAt,
			&adminData.UpdatedAt,
			&branchName,
		)
		helper.PanicIfError(err)
		return adminData, nil
	} else {
		return adminData, errors.New("admin not found")
	}
}

func (adminRepository *AdminRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, admin *domain.Admin) {
	SQL := "INSERT INTO admins (id_address, id_branch, code, level, full_name, password, email, telephone_number, photo, status) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		admin.IdAddress,
		admin.IdBranch,
		admin.Code,
		admin.Level,
		admin.FullName,
		admin.Password,
		admin.Email,
		admin.TelephoneNumber,
		admin.Photo,
		admin.Status,
	)
	helper.PanicIfError(err)
	newAdminId, err := result.LastInsertId()
	helper.PanicIfError(err)

	admin.Id = int(newAdminId)
}

func (adminRepository *AdminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin *domain.Admin) {
	SQL := "UPDATE admins SET id_branch = ?, code = ?, level = ?, full_name = ?, password = ?, telephone_number = ?, photo = ?, status = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		admin.IdBranch,
		admin.Code,
		admin.Level,
		admin.FullName,
		admin.Password,
		admin.TelephoneNumber,
		admin.Photo,
		admin.Status,
		admin.UpdatedAt,
		admin.Id,
	)
	helper.PanicIfError(err)
}

func (adminRepository *AdminRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, adminId int) {
	SQL := "DELETE FROM admins WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, adminId)
	helper.PanicIfError(err)
}
