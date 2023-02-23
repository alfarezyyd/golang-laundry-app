package usecase

import (
	"context"
	"golang-laundry-app/model/web/inventory"
	"golang-laundry-app/model/web/response"
)

type InventoryUsecase interface {
	FindAll(ctx context.Context) []response.InventoryResponse
	FindById(ctx context.Context, inventoryId int) response.InventoryResponse
	Create(ctx context.Context, inventoryCreateRequest *inventory.CreateRequestInventory) response.InventoryResponse
	Update(ctx context.Context, inventoryUpdateRequest *inventory.UpdateRequestInventory) response.InventoryResponse
	Delete(ctx context.Context, inventoryId int)
	FindAllInventoryByOrder(ctx context.Context, orderId *int) []*response.InventoryResponse
}
