package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type PromoRepositoryImpl struct {
}

func NewPromoRepositoryImpl() *PromoRepositoryImpl {
	return &PromoRepositoryImpl{}
}

func (promoRepository *PromoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Promo {
	SQL := "SELECT id, code, name, discount FROM promos"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allPromoData []domain.Promo
	for rows.Next() {
		promoData := domain.Promo{}
		err = rows.Scan(
			&promoData.Id,
			&promoData.Code,
			&promoData.Name,
			&promoData.Discount,
		)
		helper.PanicIfError(err)
		allPromoData = append(allPromoData, promoData)
	}
	return allPromoData
}

func (promoRepository *PromoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, promoId int) (domain.Promo, error) {
	SQL := "SELECT id, code, name, discount, description, status, photo, start, end, created_at, updated_at FROM promos WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, promoId)
	defer rows.Close()
	helper.PanicIfError(err)

	var promoData domain.Promo
	if rows.Next() {
		err = rows.Scan(
			&promoData.Id,
			&promoData.Code,
			&promoData.Name,
			&promoData.Discount,
			&promoData.Description,
			&promoData.Status,
			&promoData.Photo,
			&promoData.Start,
			&promoData.End,
			&promoData.CreatedAt,
			&promoData.UpdatedAt,
		)
		helper.PanicIfError(err)
		return promoData, nil
	} else {
		return promoData, errors.New("promo not found")
	}
}

func (promoRepository *PromoRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, promo *domain.Promo) {
	SQL := "INSERT INTO promos  (code, name, discount, description, status, photo, start, end) VALUES (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		promo.Code,
		promo.Name,
		promo.Discount,
		promo.Description,
		promo.Status,
		promo.Photo,
		promo.Start,
		promo.End,
	)
	helper.PanicIfError(err)
	newPromoId, err := result.LastInsertId()
	helper.PanicIfError(err)

	promo.Id = int(newPromoId)
}

func (promoRepository *PromoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, promo *domain.Promo) {
	SQL := "UPDATE promos SET code = ?, name = ?, discount = ?, description = ?, status = ?, photo = ?, start = ?, end = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		promo.Code,
		promo.Name,
		promo.Discount,
		promo.Description,
		promo.Status,
		promo.Photo,
		promo.Start,
		promo.End,
		promo.UpdatedAt,
		promo.Id,
	)
	helper.PanicIfError(err)
}

func (promoRepository *PromoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, promoId int) {
	SQL := "DELETE FROM promos WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, promoId)
	helper.PanicIfError(err)
}

func (promoRepository *PromoRepositoryImpl) FindAllPromoByOrder(ctx context.Context, tx *sql.Tx, orderId *int) []domain.Promo {
	SQL := "SELECT p.id, p.code, p.name, p.discount FROM promos AS p JOIN orders_promos AS op ON p.id = op.id_promo JOIN orders AS ord ON op.id_order = ord.id WHERE id_order = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	defer rows.Close()
	helper.PanicIfError(err)

	var allPromoData []domain.Promo
	for rows.Next() {
		promoData := domain.Promo{}
		err = rows.Scan(
			&promoData.Id,
			&promoData.Code,
			&promoData.Name,
			&promoData.Discount,
		)
		helper.PanicIfError(err)
		allPromoData = append(allPromoData, promoData)
	}
	return allPromoData
}

func (promoRepository *PromoRepositoryImpl) DeletePromoOrder(ctx context.Context, tx *sql.Tx, promoId *int) {
	SQL := "DELETE FROM orders_promos WHERE id_promo = ?"
	_, err := tx.ExecContext(ctx, SQL, promoId)
	helper.PanicIfError(err)

}
