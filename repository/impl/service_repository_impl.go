package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type ServiceRepositoryImpl struct {
}

func NewServiceRepositoryImpl() *ServiceRepositoryImpl {
	return &ServiceRepositoryImpl{}
}

func (serviceRepository *ServiceRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Service {
	SQL := "SELECT id, code, name, price, duration, created_at, updated_at FROM services"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)

	var allServiceData []domain.Service
	for rows.Next() {
		serviceData := domain.Service{}
		err = rows.Scan(
			&serviceData.Id,
			&serviceData.Code,
			&serviceData.Name,
			&serviceData.Price,
			&serviceData.Duration,
			&serviceData.CreatedAt,
			&serviceData.UpdatedAt,
		)
		helper.PanicIfError(err)
		allServiceData = append(allServiceData, serviceData)
	}
	return allServiceData
}

func (serviceRepository *ServiceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, serviceId int) (domain.Service, error) {
	SQL := "SELECT id, code, name, price, duration, created_at, updated_at FROM services WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, serviceId)
	defer rows.Close()
	helper.PanicIfError(err)

	var serviceData domain.Service
	if rows.Next() {
		err = rows.Scan(
			&serviceData.Id,
			&serviceData.Code,
			&serviceData.Name,
			&serviceData.Price,
			&serviceData.Duration,
			&serviceData.CreatedAt,
			&serviceData.UpdatedAt,
		)
		helper.PanicIfError(err)
		return serviceData, nil
	} else {
		return serviceData, errors.New("service not found")
	}
}

func (serviceRepository *ServiceRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, service *domain.Service) {
	SQL := "INSERT INTO services(code, name, price, duration) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		service.Code,
		service.Name,
		service.Price,
		service.Duration,
	)
	helper.PanicIfError(err)

	newServiceId, err := result.LastInsertId()
	helper.PanicIfError(err)

	service.Id = int(newServiceId)
}

func (serviceRepository *ServiceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, service *domain.Service) {
	SQL := "UPDATE services SET code = ?, name = ?, price = ?, duration = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		service.Code,
		service.Name,
		service.Price,
		service.Duration,
		service.UpdatedAt,
		service.Id,
	)
	helper.PanicIfError(err)
}
func (serviceRepository *ServiceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, serviceId int) {
	SQL := "DELETE FROM services WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, serviceId)
	helper.PanicIfError(err)
}
