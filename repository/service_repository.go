package repository

import (
	"context"
	"database/sql"
	"golang-laundry-app/model/domain"
)

type ServiceRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Service
	FindById(ctx context.Context, tx *sql.Tx, serviceId int) (domain.Service, error)
	Create(ctx context.Context, tx *sql.Tx, service *domain.Service)
	Update(ctx context.Context, tx *sql.Tx, service *domain.Service)
	Delete(ctx context.Context, tx *sql.Tx, serviceId int)
}
