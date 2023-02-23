package usecase

import (
	"context"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/model/web/service"
)

type ServiceUsecase interface {
	FindAll(ctx context.Context) []response.ServiceResponse
	FindById(ctx context.Context, serviceId int) response.ServiceResponse
	Create(ctx context.Context, serviceCreateRequest *service.CreateRequestService) response.ServiceResponse
	Update(ctx context.Context, serviceUpdateRequest *service.UpdateRequestService) response.ServiceResponse
	Delete(ctx context.Context, serviceId int)
}
