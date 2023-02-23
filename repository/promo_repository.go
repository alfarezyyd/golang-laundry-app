package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type PromoRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Promo
	FindById(ctx context.Context, tx *sql.Tx, promoId int) (domain.Promo, error)
	Create(ctx context.Context, tx *sql.Tx, promo *domain.Promo)
	Update(ctx context.Context, tx *sql.Tx, promo *domain.Promo)
	Delete(ctx context.Context, tx *sql.Tx, promoId int)
	FindAllPromoByOrder(ctx context.Context, tx *sql.Tx, orderId *int) []domain.Promo
	DeletePromoOrder(ctx context.Context, tx *sql.Tx, promoId *int)
}
