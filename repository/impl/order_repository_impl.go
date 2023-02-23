package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (orderRepository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, allEmployeeName, allUserName, allServiceName []*string) []domain.Order {
	SQL := "SELECT o.id, o.code, o.price, o.status, users.full_name, employees.full_name, services.name FROM orders AS o JOIN users ON o.id_user = users.id JOIN employees ON o.id_employee = employees.id JOIN services ON o.id_service = services.id"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allOrderData []domain.Order
	for rows.Next() {
		var userName, employeeName, serviceName *string
		orderData := domain.Order{}
		err = rows.Scan(
			&orderData.Id,
			&orderData.Code,
			&orderData.Price,
			&orderData.Status,
			userName,
			employeeName,
			serviceName,
		)
		helper.PanicIfError(err)
		allOrderData = append(allOrderData, orderData)
		allUserName = append(allUserName, userName)
		allEmployeeName = append(allEmployeeName, employeeName)
		allServiceName = append(allServiceName, serviceName)
	}
	return allOrderData
}

func (orderRepository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int, employeeName, userName, serviceName string) (domain.Order, error) {
	SQL := "SELECT o.id, o.id_user, o.id_employee, o.id_service, o.code, o.type, o.price, o.weight, o.payment, o.description, o.status, o.entry, o.out, o.created_at, o.updated_at, users.full_name, employees.full_name, services.name FROM orders AS o JOIN users ON o.id_user = users.id JOIN employees ON o.id_employee = employees.id JOIN services ON o.id_service = services.id WHERE o.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	defer rows.Close()
	helper.PanicIfError(err)

	var orderData domain.Order
	if rows.Next() {
		err = rows.Scan(
			&orderData.Id,
			&orderData.IdUser,
			&orderData.IdEmployee,
			&orderData.IdService,
			&orderData.Code,
			&orderData.Type,
			&orderData.Price,
			&orderData.Weight,
			&orderData.Payment,
			&orderData.Description,
			&orderData.Status,
			&orderData.Entry,
			&orderData.Out,
			&orderData.CreatedAt,
			&orderData.UpdatedAt,
			&userName,
			&employeeName,
			&serviceName,
		)
		helper.PanicIfError(err)
		return orderData, nil
	} else {
		return orderData, errors.New("order not found")
	}
}

func (orderRepository *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order *domain.Order) {
	SQL := "INSERT INTO orders (id_user, id_employee, id_service, code, type, price, weight, payment, description, status, entry, `out`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		order.IdUser,
		order.IdEmployee,
		order.IdService,
		order.Code,
		order.Type,
		order.Price,
		order.Weight,
		order.Payment,
		order.Description,
		order.Status,
		order.Entry,
		order.Out,
	)
	helper.PanicIfError(err)
	newOrderId, err := result.LastInsertId()
	helper.PanicIfError(err)
	order.Id = int(newOrderId)
}

func (orderRepository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order *domain.Order) {
	SQL := "UPDATE orders SET id_user = ?, id_employee = ?, id_service = ?, type = ?, price = ?, weight = ?, payment = ?, description = ?, status = ?, entry = ?, `out` = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		order.IdUser,
		order.IdEmployee,
		order.IdService,
		order.Type,
		order.Price,
		order.Weight,
		order.Payment,
		order.Description,
		order.Status,
		order.Entry,
		order.Out,
		order.UpdatedAt,
		order.Id,
	)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderId int) {
	SQL := "DELETE FROM orders WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) FindAllOrderByUser(ctx context.Context, tx *sql.Tx, userId *int) []int {
	SQL := "SELECT id FROM orders WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	defer rows.Close()
	helper.PanicIfError(err)

	var allIdOrder []int
	for rows.Next() {
		var idOrder int
		err = rows.Scan(&idOrder)
		helper.PanicIfError(err)
		allIdOrder = append(allIdOrder, idOrder)
	}
	return allIdOrder
}

func (orderRepository *OrderRepositoryImpl) CreateOrderInventory(ctx context.Context, tx *sql.Tx, orderId, inventoryId *int) {
	SQL := "INSERT INTO orders_inventories (id_order, id_inventory) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, orderId, inventoryId)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) CreateOrderPromo(ctx context.Context, tx *sql.Tx, orderId, promoId *int) {
	SQL := "INSERT INTO orders_promos (id_order, id_promo) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, orderId, promoId)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) DeleteOrderInventory(ctx context.Context, tx *sql.Tx, orderId *int) {
	SQL := "DELETE FROM orders_inventories WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) DeleteOrderPromo(ctx context.Context, tx *sql.Tx, orderId *int) {
	SQL := "DELETE FROM orders_promos WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
}

func (orderRepository *OrderRepositoryImpl) DeleteAllOrderByUser(ctx context.Context, tx *sql.Tx, userId *int) {
	SQL := "DELETE FROM orders WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}
