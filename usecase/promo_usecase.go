package usecase

import (
	"context"
	"golang-laundry-app/model/web/promo"
	"golang-laundry-app/model/web/response"
)

type PromoUsecase interface {
	FindAll(ctx context.Context) []response.PromoResponse
	FindById(ctx context.Context, promoId int) response.PromoResponse
	Create(ctx context.Context, promoCreateRequest *promo.CreateRequestPromo) response.PromoResponse
	Update(ctx context.Context, promoUpdateRequest *promo.UpdateRequestPromo) response.PromoResponse
	Delete(ctx context.Context, promoId int)
	FindAllPromoByOrder(ctx context.Context, orderId *int) []*response.PromoResponse
}
