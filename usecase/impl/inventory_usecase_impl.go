package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/inventory"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"time"
)

type InventoryUsecaseImpl struct {
	InventoryRepository repository.InventoryRepository
	EmployeeRepository  repository.EmployeeRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewInventoryUsecaseImpl(inventoryRepository repository.InventoryRepository, employeeRepository repository.EmployeeRepository, DB *sql.DB, validate *validator.Validate) *InventoryUsecaseImpl {
	return &InventoryUsecaseImpl{InventoryRepository: inventoryRepository, EmployeeRepository: employeeRepository, DB: DB, Validate: validate}
}

func (inventoryUsecase *InventoryUsecaseImpl) FindAll(ctx context.Context) []response.InventoryResponse {
	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allInventoryData := inventoryUsecase.InventoryRepository.FindAll(ctx, tx)
	var allInventoryResponse []response.InventoryResponse
	for _, inventoryData := range allInventoryData {
		inventoryResponse := response.InventoryResponse{
			Id:        inventoryData.Id,
			Code:      inventoryData.Code,
			Commodity: inventoryData.Commodity,
			Variant:   inventoryData.Variant,
			Price:     inventoryData.Price,
		}
		allInventoryResponse = append(allInventoryResponse, inventoryResponse)
	}
	return allInventoryResponse
}

func (inventoryUsecase *InventoryUsecaseImpl) FindById(ctx context.Context, inventoryId int) response.InventoryResponse {
	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	inventoryData, err := inventoryUsecase.InventoryRepository.FindById(ctx, tx, inventoryId)
	exception.ResponseIfNotFoundError(err)

	employeeData, err := inventoryUsecase.EmployeeRepository.FindById(ctx, tx, inventoryData.IdEmployee, "")
	exception.ResponseIfNotFoundError(err)
	return helper.ConvertToInventoryResponse(&inventoryData, &employeeData)
}

func (inventoryUsecase *InventoryUsecaseImpl) Create(ctx context.Context, inventoryCreateRequest *inventory.CreateRequestInventory) response.InventoryResponse {
	err := inventoryUsecase.Validate.Struct(inventoryCreateRequest)
	helper.PanicIfError(err)

	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	inventoryData := domain.Inventory{
		IdEmployee: inventoryCreateRequest.IdEmployee,
		Code:       inventoryCreateRequest.Code,
		Commodity:  inventoryCreateRequest.Commodity,
		Variant:    inventoryCreateRequest.Variant,
		Quantity:   inventoryCreateRequest.Quantity,
		Price:      inventoryCreateRequest.Price,
		Supplier:   inventoryCreateRequest.Supplier,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	inventoryUsecase.InventoryRepository.Create(ctx, tx, &inventoryData)

	employeeData, err := inventoryUsecase.EmployeeRepository.FindById(ctx, tx, inventoryData.IdEmployee, "")
	exception.ResponseIfNotFoundError(err)

	return helper.ConvertToInventoryResponse(&inventoryData, &employeeData)
}

func (inventoryUsecase *InventoryUsecaseImpl) Update(ctx context.Context, inventoryUpdateRequest *inventory.UpdateRequestInventory) response.InventoryResponse {
	err := inventoryUsecase.Validate.Struct(inventoryUpdateRequest)
	helper.PanicIfError(err)

	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	inventoryData, err := inventoryUsecase.InventoryRepository.FindById(ctx, tx, inventoryUpdateRequest.Id)
	exception.ResponseIfNotFoundError(err)

	inventoryData.IdEmployee = inventoryUpdateRequest.IdEmployee
	inventoryData.Code = inventoryUpdateRequest.Code
	inventoryData.Commodity = inventoryUpdateRequest.Commodity
	inventoryData.Variant = inventoryUpdateRequest.Variant
	inventoryData.Quantity = inventoryUpdateRequest.Quantity
	inventoryData.Price = inventoryUpdateRequest.Price
	inventoryData.Supplier = inventoryUpdateRequest.Supplier
	inventoryData.UpdatedAt = time.Now()

	inventoryUsecase.InventoryRepository.Update(ctx, tx, &inventoryData)

	employeeData, err := inventoryUsecase.EmployeeRepository.FindById(ctx, tx, inventoryData.IdEmployee, "")
	exception.ResponseIfNotFoundError(err)

	return helper.ConvertToInventoryResponse(&inventoryData, &employeeData)

}

func (inventoryUsecase *InventoryUsecaseImpl) Delete(ctx context.Context, inventoryId int) {
	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	_, err = inventoryUsecase.InventoryRepository.FindById(ctx, tx, inventoryId)
	exception.ResponseIfNotFoundError(err)

	inventoryUsecase.InventoryRepository.Delete(ctx, tx, inventoryId)
}

func (inventoryUsecase *InventoryUsecaseImpl) FindAllInventoryByOrder(ctx context.Context, orderId *int) []*response.InventoryResponse {
	tx, err := inventoryUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allInventoryData := inventoryUsecase.InventoryRepository.FindAllInventoryByOrder(ctx, tx, orderId)
	var allInventoryResponse []*response.InventoryResponse
	for _, inventoryData := range allInventoryData {
		inventoryResponse := response.InventoryResponse{
			Id:        inventoryData.Id,
			Code:      inventoryData.Code,
			Commodity: inventoryData.Commodity,
			Variant:   inventoryData.Variant,
			Price:     inventoryData.Price,
		}
		allInventoryResponse = append(allInventoryResponse, &inventoryResponse)
	}
	return allInventoryResponse
}
