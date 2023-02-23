package usecase

import (
	"context"
	"golang-laundry-app/model/web/order"
	"golang-laundry-app/model/web/response"
)

type OrderUsecase interface {
	FindAll(ctx context.Context) []response.OrderResponse
	FindById(ctx context.Context, orderId int) response.OrderResponse
	Create(ctx context.Context, orderCreateRequest *order.CreateRequestOrder) response.OrderResponse
	Update(ctx context.Context, orderUpdateRequest *order.UpdateRequestOrder) response.OrderResponse
	Delete(ctx context.Context, orderId int)
	DeleteAllOrderByUser(ctx context.Context, userId *int)
}
