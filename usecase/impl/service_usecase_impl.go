package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/model/web/service"
	"golang-laundry-app/repository"
	"time"
)

type ServiceUsecaseImpl struct {
	ServiceRepository repository.ServiceRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewServiceUsecaseImpl(serviceRepository repository.ServiceRepository, DB *sql.DB, validate *validator.Validate) *ServiceUsecaseImpl {
	return &ServiceUsecaseImpl{ServiceRepository: serviceRepository, DB: DB, Validate: validate}
}

func (serviceUsecase *ServiceUsecaseImpl) FindAll(ctx context.Context) []response.ServiceResponse {
	tx, err := serviceUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allServiceData := serviceUsecase.ServiceRepository.FindAll(ctx, tx)
	var allServiceResponse []response.ServiceResponse
	for _, serviceData := range allServiceData {
		serviceResponse := response.ServiceResponse{
			Id:        serviceData.Id,
			Code:      serviceData.Code,
			Name:      serviceData.Name,
			Price:     serviceData.Price,
			Duration:  serviceData.Duration,
			CreatedAt: serviceData.CreatedAt,
			UpdatedAt: serviceData.UpdatedAt,
		}
		allServiceResponse = append(allServiceResponse, serviceResponse)
	}
	return allServiceResponse
}

func (serviceUsecase *ServiceUsecaseImpl) FindById(ctx context.Context, serviceId int) response.ServiceResponse {
	tx, err := serviceUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	serviceData, err := serviceUsecase.ServiceRepository.FindById(ctx, tx, serviceId)
	exception.ResponseIfNotFoundError(err)

	return helper.ConvertToServiceResponse(&serviceData)
}

func (serviceUsecase *ServiceUsecaseImpl) Create(ctx context.Context, serviceCreateRequest *service.CreateRequestService) response.ServiceResponse {
	err := serviceUsecase.Validate.Struct(serviceCreateRequest)
	helper.PanicIfError(err)

	tx, err := serviceUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	serviceData := domain.Service{
		Code:      serviceCreateRequest.Code,
		Name:      serviceCreateRequest.Name,
		Price:     serviceCreateRequest.Price,
		Duration:  serviceCreateRequest.Duration,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	serviceUsecase.ServiceRepository.Create(ctx, tx, &serviceData)
	return helper.ConvertToServiceResponse(&serviceData)
}

func (serviceUsecase *ServiceUsecaseImpl) Update(ctx context.Context, serviceUpdateRequest *service.UpdateRequestService) response.ServiceResponse {
	err := serviceUsecase.Validate.Struct(serviceUpdateRequest)
	helper.PanicIfError(err)

	tx, err := serviceUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	serviceData, err := serviceUsecase.ServiceRepository.FindById(ctx, tx, serviceUpdateRequest.Id)
	exception.ResponseIfNotFoundError(err)

	serviceData.Id = serviceUpdateRequest.Id
	serviceData.Code = serviceUpdateRequest.Code
	serviceData.Name = serviceUpdateRequest.Name
	serviceData.Price = serviceUpdateRequest.Price
	serviceData.Duration = serviceUpdateRequest.Duration
	serviceData.UpdatedAt = time.Now()

	serviceUsecase.ServiceRepository.Update(ctx, tx, &serviceData)
	return helper.ConvertToServiceResponse(&serviceData)
}

func (serviceUsecase *ServiceUsecaseImpl) Delete(ctx context.Context, serviceId int) {
	tx, err := serviceUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	_, err = serviceUsecase.ServiceRepository.FindById(ctx, tx, serviceId)
	exception.ResponseIfNotFoundError(err)
	serviceUsecase.ServiceRepository.Delete(ctx, tx, serviceId)

}
