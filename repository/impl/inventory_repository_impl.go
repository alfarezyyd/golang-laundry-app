package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type InventoryRepositoryImpl struct {
}

func NewInventoryRepositoryImpl() *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{}
}

func (inventoryRepository *InventoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Inventory {
	SQL := "SELECT id, code, commodity, variant, price FROM inventories"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allInventoryData []domain.Inventory
	for rows.Next() {
		inventoryData := domain.Inventory{}
		err = rows.Scan(
			&inventoryData.Id,
			&inventoryData.Code,
			&inventoryData.Commodity,
			&inventoryData.Variant,
			&inventoryData.Price,
		)
		helper.PanicIfError(err)
		allInventoryData = append(allInventoryData, inventoryData)
	}
	return allInventoryData
}

func (inventoryRepository *InventoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, inventoryId int) (domain.Inventory, error) {
	SQL := "SELECT id, id_employee, code, commodity, variant, quantity, price, supplier, created_at, updated_at FROM inventories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, inventoryId)
	defer rows.Close()
	helper.PanicIfError(err)

	var inventoryData domain.Inventory
	if rows.Next() {
		err = rows.Scan(
			&inventoryData.Id,
			&inventoryData.IdEmployee,
			&inventoryData.Code,
			&inventoryData.Commodity,
			&inventoryData.Variant,
			&inventoryData.Quantity,
			&inventoryData.Price,
			&inventoryData.Supplier,
			&inventoryData.CreatedAt,
			&inventoryData.UpdatedAt,
		)
		helper.PanicIfError(err)
		return inventoryData, nil
	} else {
		return inventoryData, errors.New("inventory not found")
	}
}

func (inventoryRepository *InventoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, inventory *domain.Inventory) {
	SQL := "INSERT INTO inventories(id_employee, code, commodity, variant, quantity, price, supplier, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		inventory.IdEmployee,
		inventory.Code,
		inventory.Commodity,
		inventory.Variant,
		inventory.Quantity,
		inventory.Price,
		inventory.Supplier,
		inventory.CreatedAt,
		inventory.UpdatedAt,
	)
	newInventoryId, err := result.LastInsertId()
	helper.PanicIfError(err)

	inventory.Id = int(newInventoryId)
}

func (inventoryRepository *InventoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, inventory *domain.Inventory) {
	SQL := "UPDATE inventories SET id_employee = ?, code = ?, commodity= ? , variant = ?, quantity = ?, price = ?, supplier = ?,  updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		inventory.IdEmployee,
		inventory.Code,
		inventory.Commodity,
		inventory.Variant,
		inventory.Quantity,
		inventory.Price,
		inventory.Supplier,
		inventory.UpdatedAt,
		inventory.Id,
	)
	helper.PanicIfError(err)
}

func (inventoryRepository *InventoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, inventoryId int) {
	SQL := "DELETE FROM inventories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, inventoryId)
	helper.PanicIfError(err)
}

func (inventoryRepository *InventoryRepositoryImpl) FindAllInventoryByOrder(ctx context.Context, tx *sql.Tx, orderId *int) []domain.Inventory {
	SQL := "SELECT inv.id, inv.code, inv.commodity, inv.variant, inv.price FROM inventories AS inv JOIN orders_inventories AS ord ON ord.id_inventory = inv.id JOIN orders AS o ON ord.id_order = o.id WHERE id_order = ?;"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	defer rows.Close()
	helper.PanicIfError(err)

	var allInventoryData []domain.Inventory
	for rows.Next() {
		inventoryData := domain.Inventory{}
		err = rows.Scan(
			&inventoryData.Id,
			&inventoryData.Code,
			&inventoryData.Commodity,
			&inventoryData.Variant,
			&inventoryData.Price,
		)
		helper.PanicIfError(err)
		allInventoryData = append(allInventoryData, inventoryData)
	}
	return allInventoryData
}
