package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, level, full_name, gender FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allUserData []domain.User
	for rows.Next() {
		userData := domain.User{}
		err = rows.Scan(
			&userData.Id,
			&userData.Level,
			&userData.FullName,
			&userData.Gender,
		)
		helper.PanicIfError(err)
		allUserData = append(allUserData, userData)
	}
	return allUserData
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, id_address, level, full_name, gender, password, email, telephone_number, photo, email_verified_at, created_at, updated_at FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	defer rows.Close()
	helper.PanicIfError(err)

	var userData domain.User
	if rows.Next() {
		err = rows.Scan(
			&userData.Id,
			&userData.IdAddress,
			&userData.Level,
			&userData.FullName,
			&userData.Gender,
			&userData.Password,
			&userData.Email,
			&userData.TelephoneNumber,
			&userData.Photo,
			&userData.EmailVerifiedAt,
			&userData.CreatedAt,
			&userData.UpdatedAt,
		)
		helper.PanicIfError(err)
		return userData, nil
	} else {
		return userData, errors.New("user not found")
	}
}

func (userRepository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user *domain.User) {
	SQL := "INSERT INTO users (id_address, level, full_name, gender, password, email, telephone_number, photo) VALUES (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		user.IdAddress,
		user.Level,
		user.FullName,
		user.Gender,
		user.Password,
		user.Email,
		user.TelephoneNumber,
		user.Photo,
	)
	helper.PanicIfError(err)
	newUserId, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(newUserId)
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) {
	SQL := "UPDATE users SET level = ?, full_name = ?, gender = ?, password = ?, telephone_number = ?, photo = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Level,
		user.FullName,
		user.Gender,
		user.Password,
		user.TelephoneNumber,
		user.Photo,
		user.UpdatedAt,
		user.Id,
	)
	helper.PanicIfError(err)
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

func (userRepository *UserRepositoryImpl) FindAuthUser(ctx context.Context, tx *sql.Tx, userEmail, userPassword string) (domain.User, error) {
	SQL := "SELECT email, password FROM users WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, userEmail, userPassword)
	defer rows.Close()
	helper.PanicIfError(err)

	var userAuthData domain.User
	if rows.Next() {
		err = rows.Scan(&userAuthData.Email, &userAuthData.Password)
		helper.PanicIfError(err)
		return userAuthData, nil
	} else {
		return userAuthData, errors.New("user not found")
	}
}
