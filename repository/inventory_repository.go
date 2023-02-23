package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type InventoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Inventory
	FindById(ctx context.Context, tx *sql.Tx, inventoryId int) (domain.Inventory, error)
	Create(ctx context.Context, tx *sql.Tx, inventory *domain.Inventory)
	Update(ctx context.Context, tx *sql.Tx, inventory *domain.Inventory)
	Delete(ctx context.Context, tx *sql.Tx, inventoryId int)
	FindAllInventoryByOrder(ctx context.Context, tx *sql.Tx, orderId *int) []domain.Inventory
}
