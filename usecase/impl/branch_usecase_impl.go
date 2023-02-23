package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/branch"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"golang-laundry-app/usecase"
	"time"
)

type BranchUsecaseImpl struct {
	BranchRepository repository.BranchRepository
	AddressUsecase   usecase.AddressUsecase
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewBranchUsecaseImpl(branchRepository repository.BranchRepository, addressUsecase usecase.AddressUsecase, DB *sql.DB, validate *validator.Validate) *BranchUsecaseImpl {
	return &BranchUsecaseImpl{BranchRepository: branchRepository, AddressUsecase: addressUsecase, DB: DB, Validate: validate}
}

func (branchUsecase *BranchUsecaseImpl) FindAll(ctx context.Context) []response.BranchResponse {
	tx, err := branchUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allBranchData := branchUsecase.BranchRepository.FindAll(ctx, tx)
	var allBranchResponse []response.BranchResponse
	for _, branchData := range allBranchData {
		branchResponse := response.BranchResponse{
			Id:              branchData.Id,
			Name:            branchData.Name,
			TelephoneNumber: branchData.TelephoneNumber,
			CreatedAt:       &branchData.CreatedAt,
			UpdatedAt:       &branchData.UpdatedAt,
		}
		allBranchResponse = append(allBranchResponse, branchResponse)
	}
	return allBranchResponse
}

func (branchUsecase *BranchUsecaseImpl) FindById(ctx context.Context, branchId int) response.BranchResponse {
	tx, err := branchUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	branchData, err := branchUsecase.BranchRepository.FindById(ctx, tx, branchId)
	exception.ResponseIfNotFoundError(err)

	addressData := branchUsecase.AddressUsecase.FindById(ctx, branchData.IdAddress)
	return helper.ConvertToBranchResponse(&branchData, &addressData)
}

func (branchUsecase *BranchUsecaseImpl) Create(ctx context.Context, branchCreateRequest *branch.CreateRequestBranch) response.BranchResponse {
	err := branchUsecase.Validate.Struct(branchCreateRequest)
	helper.PanicIfError(err)

	tx, err := branchUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressData := branchUsecase.AddressUsecase.Create(ctx, branchCreateRequest.Address)

	branchData := domain.Branch{
		IdAddress:       addressData.Id,
		Name:            branchCreateRequest.Name,
		TelephoneNumber: branchCreateRequest.TelephoneNumber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	branchUsecase.BranchRepository.Create(ctx, tx, &branchData)

	return helper.ConvertToBranchResponse(&branchData, &addressData)
}

func (branchUsecase *BranchUsecaseImpl) Update(ctx context.Context, branchUpdateRequest *branch.UpdateRequestBranch) response.BranchResponse {
	err := branchUsecase.Validate.Struct(branchUpdateRequest)
	helper.PanicIfError(err)

	tx, err := branchUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	_, err = branchUsecase.BranchRepository.FindById(ctx, tx, branchUpdateRequest.Id)
	exception.ResponseIfNotFoundError(err)

	addressData := branchUsecase.AddressUsecase.Update(ctx, branchUpdateRequest.Address)

	branchData := domain.Branch{
		Id:              branchUpdateRequest.Id,
		Name:            branchUpdateRequest.Name,
		TelephoneNumber: branchUpdateRequest.TelephoneNumber,
		UpdatedAt:       time.Now(),
	}
	branchUsecase.BranchRepository.Update(ctx, tx, &branchData)
	return helper.ConvertToBranchResponse(&branchData, &addressData)

}

func (branchUsecase *BranchUsecaseImpl) Delete(ctx context.Context, branchId int) {
	tx, err := branchUsecase.DB.Begin()
	helper.PanicIfError(err)

	branchData := branchUsecase.FindById(ctx, branchId)
	branchUsecase.BranchRepository.Delete(ctx, tx, branchData.Id)
	helper.CommitOrRollback(tx)

	branchUsecase.AddressUsecase.Delete(ctx, branchData.Address.Id)
}
