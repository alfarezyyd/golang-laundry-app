package usecase

import (
	"context"
	"golang-laundry-app/model/web/admin"
	"golang-laundry-app/model/web/response"
)

type AdminUsecase interface {
	FindAll(ctx context.Context) []response.AdminResponse
	FindById(ctx context.Context, adminId int) response.AdminResponse
	Create(ctx context.Context, createRequestAdmin *admin.CreateRequestAdmin) response.AdminResponse
	Update(ctx context.Context, updateRequestAdmin *admin.UpdateRequestAdmin) response.AdminResponse
	Delete(ctx context.Context, adminId int)
}
