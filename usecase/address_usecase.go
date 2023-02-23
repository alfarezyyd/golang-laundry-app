package usecase

import (
	"context"
	"golang-laundry-app/model/web/address"
	"golang-laundry-app/model/web/response"
)

type AddressUsecase interface {
	FindById(ctx context.Context, addressId int) response.AddressResponse
	Create(ctx context.Context, addressCreateRequest *address.CreateRequestAddress) response.AddressResponse
	Update(ctx context.Context, addressUpdateRequest *address.UpdateRequestAddress) response.AddressResponse
	Delete(ctx context.Context, addressId int)
}
