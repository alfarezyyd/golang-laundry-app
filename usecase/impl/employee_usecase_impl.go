package impl

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/employee"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"golang-laundry-app/usecase"
	"time"
)

type EmployeeUsecaseImpl struct {
	EmployeeRepository repository.EmployeeRepository
	BranchUsecase      usecase.BranchUsecase
	AddressUsecase     usecase.AddressUsecase
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewEmployeeUsecaseImpl(employeeRepository repository.EmployeeRepository, branchUsecase usecase.BranchUsecase, addressUsecase usecase.AddressUsecase, DB *sql.DB, validate *validator.Validate) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{EmployeeRepository: employeeRepository, BranchUsecase: branchUsecase, AddressUsecase: addressUsecase, DB: DB, Validate: validate}
}

func (employeeUsecase *EmployeeUsecaseImpl) FindAll(ctx context.Context) []response.EmployeeResponse {
	tx, err := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allEmployeeData := employeeUsecase.EmployeeRepository.FindAll(ctx, tx)
	var allEmployeeResponse []response.EmployeeResponse
	for _, employeeData := range allEmployeeData {
		employeeResponse := response.EmployeeResponse{
			Id:       employeeData.Id,
			Code:     employeeData.Code,
			Level:    employeeData.Level,
			FullName: employeeData.FullName,
			Status:   employeeData.Status,
		}
		allEmployeeResponse = append(allEmployeeResponse, employeeResponse)
	}
	return allEmployeeResponse
}

func (employeeUsecase *EmployeeUsecaseImpl) FindById(ctx context.Context, employeeId int) response.EmployeeResponse {
	tx, err := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	var branchName string
	employeeDetailData, err := employeeUsecase.EmployeeRepository.FindById(ctx, tx, employeeId, branchName)
	exception.ResponseIfNotFoundError(err)

	employeeAddressData := employeeUsecase.AddressUsecase.FindById(ctx, employeeDetailData.IdAddress)

	return helper.ConvertToEmployeeResponse(&employeeDetailData, &employeeAddressData, &branchName)
}

func (employeeUsecase *EmployeeUsecaseImpl) Create(ctx context.Context, createRequestEmployee *employee.CreateRequestEmployee) response.EmployeeResponse {
	err := employeeUsecase.Validate.Struct(createRequestEmployee)
	helper.PanicIfError(err)

	tx, err := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressResponseData := employeeUsecase.AddressUsecase.Create(ctx, createRequestEmployee.Address)
	branchResponseData := employeeUsecase.BranchUsecase.FindById(ctx, createRequestEmployee.IdBranch)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(createRequestEmployee.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	employeeData := domain.Employee{
		IdAddress:       addressResponseData.Id,
		IdBranch:        branchResponseData.Id,
		Code:            createRequestEmployee.Code,
		Level:           createRequestEmployee.Level,
		FullName:        createRequestEmployee.FullName,
		Password:        passwordHashedByte,
		Email:           createRequestEmployee.Email,
		TelephoneNumber: createRequestEmployee.TelephoneNumber,
		Photo:           createRequestEmployee.Photo,
		Status:          createRequestEmployee.Status,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	employeeUsecase.EmployeeRepository.Create(ctx, tx, &employeeData)
	return helper.ConvertToEmployeeResponse(&employeeData, &addressResponseData, &branchResponseData.Name)
}

func (employeeUsecase *EmployeeUsecaseImpl) Update(ctx context.Context, updateRequestEmployee *employee.UpdateRequestEmployee) response.EmployeeResponse {
	err := employeeUsecase.Validate.Struct(updateRequestEmployee)
	helper.PanicIfError(err)

	tx, err := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	employeeData, err := employeeUsecase.EmployeeRepository.FindById(ctx, tx, updateRequestEmployee.Id, "")
	exception.ResponseIfNotFoundError(err)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(updateRequestEmployee.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	employeeData.Id = updateRequestEmployee.Id
	employeeData.IdBranch = updateRequestEmployee.IdBranch
	employeeData.Code = updateRequestEmployee.Code
	employeeData.Level = updateRequestEmployee.Level
	employeeData.FullName = updateRequestEmployee.FullName
	employeeData.Password = passwordHashedByte
	employeeData.TelephoneNumber = updateRequestEmployee.TelephoneNumber
	employeeData.Photo = updateRequestEmployee.Photo
	employeeData.Status = updateRequestEmployee.Status
	employeeData.UpdatedAt = time.Now()

	employeeUsecase.EmployeeRepository.Update(ctx, tx, &employeeData)
	return helper.ConvertToEmployeeResponse(&employeeData, nil, nil)
}

func (employeeUsecase *EmployeeUsecaseImpl) Delete(ctx context.Context, employeeId int) {
	tx, err := employeeUsecase.DB.Begin()
	helper.PanicIfError(err)

	employeeData := employeeUsecase.FindById(ctx, employeeId)
	employeeUsecase.EmployeeRepository.Delete(ctx, tx, employeeData.Id)
	helper.CommitOrRollback(tx)

	employeeUsecase.AddressUsecase.Delete(ctx, employeeData.Address.Id)
}
