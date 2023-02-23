package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type OrderRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, allEmployeeName, allUserName, allServiceName []*string) []domain.Order
	FindById(ctx context.Context, tx *sql.Tx, orderId int, employeeName, userName, serviceName string) (domain.Order, error)
	Create(ctx context.Context, tx *sql.Tx, order *domain.Order)
	Update(ctx context.Context, tx *sql.Tx, order *domain.Order)
	Delete(ctx context.Context, tx *sql.Tx, orderId int)
	FindAllOrderByUser(ctx context.Context, tx *sql.Tx, userId *int) []int
	CreateOrderInventory(ctx context.Context, tx *sql.Tx, orderId, inventoryId *int)
	CreateOrderPromo(ctx context.Context, tx *sql.Tx, orderId, promoId *int)
	DeleteOrderInventory(ctx context.Context, tx *sql.Tx, orderId *int)
	DeleteOrderPromo(ctx context.Context, tx *sql.Tx, orderId *int)
	DeleteAllOrderByUser(ctx context.Context, tx *sql.Tx, userId *int)
}
