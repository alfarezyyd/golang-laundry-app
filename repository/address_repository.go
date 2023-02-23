package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type AddressRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, addressId int) (domain.Address, error)
	Create(ctx context.Context, tx *sql.Tx, address *domain.Address)
	Update(ctx context.Context, tx *sql.Tx, address *domain.Address)
	Delete(ctx context.Context, tx *sql.Tx, addressId int)
}
