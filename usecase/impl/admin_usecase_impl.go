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
	"golang-laundry-app/model/web/admin"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"golang-laundry-app/usecase"
	"time"
)

type AdminUsecaseImpl struct {
	AdminRepository  repository.AdminRepository
	AddressUsecase   usecase.AddressUsecase
	BranchRepository repository.BranchRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewAdminUsecaseImpl(adminRepository repository.AdminRepository, addressUsecase usecase.AddressUsecase, branchRepository repository.BranchRepository, DB *sql.DB, validate *validator.Validate) *AdminUsecaseImpl {
	return &AdminUsecaseImpl{AdminRepository: adminRepository, AddressUsecase: addressUsecase, BranchRepository: branchRepository, DB: DB, Validate: validate}
}

func (adminUsecase *AdminUsecaseImpl) FindAll(ctx context.Context) []response.AdminResponse {
	tx, err := adminUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allAdminData := adminUsecase.AdminRepository.FindAll(ctx, tx)
	var allAdminResponse []response.AdminResponse
	for _, adminData := range allAdminData {
		adminResponse := response.AdminResponse{
			Id:       adminData.Id,
			Code:     adminData.Code,
			Level:    adminData.Level,
			FullName: adminData.FullName,
			Status:   adminData.Status,
		}
		allAdminResponse = append(allAdminResponse, adminResponse)
	}
	return allAdminResponse
}

func (adminUsecase *AdminUsecaseImpl) FindById(ctx context.Context, adminId int) response.AdminResponse {
	tx, err := adminUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	var branchName string
	adminData, err := adminUsecase.AdminRepository.FindById(ctx, tx, adminId, branchName)
	exception.ResponseIfNotFoundError(err)

	addressData := adminUsecase.AddressUsecase.FindById(ctx, adminData.IdAddress)

	return helper.ConvertToAdminResponse(&adminData, &addressData, &branchName)
}

func (adminUsecase *AdminUsecaseImpl) Create(ctx context.Context, createRequestAdmin *admin.CreateRequestAdmin) response.AdminResponse {
	err := adminUsecase.Validate.Struct(createRequestAdmin)
	helper.PanicIfError(err)

	tx, err := adminUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(createRequestAdmin.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	addressData := adminUsecase.AddressUsecase.Create(ctx, createRequestAdmin.Address)
	branchResponseData, err := adminUsecase.BranchRepository.FindById(ctx, tx, createRequestAdmin.IdBranch)
	exception.ResponseIfNotFoundError(err)

	adminData := domain.Admin{
		IdAddress:       addressData.Id,
		IdBranch:        branchResponseData.Id,
		Code:            createRequestAdmin.Code,
		Level:           createRequestAdmin.Level,
		FullName:        createRequestAdmin.FullName,
		Password:        passwordHashedByte,
		Email:           createRequestAdmin.Email,
		TelephoneNumber: createRequestAdmin.TelephoneNumber,
		Photo:           createRequestAdmin.Photo,
		Status:          createRequestAdmin.Status,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	adminUsecase.AdminRepository.Create(ctx, tx, &adminData)
	return helper.ConvertToAdminResponse(&adminData, &addressData, &branchResponseData.Name)
}

func (adminUsecase *AdminUsecaseImpl) Update(ctx context.Context, updateRequestAdmin *admin.UpdateRequestAdmin) response.AdminResponse {
	err := adminUsecase.Validate.Struct(updateRequestAdmin)
	helper.PanicIfError(err)

	tx, err := adminUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	branchData, err := adminUsecase.BranchRepository.FindById(ctx, tx, updateRequestAdmin.IdBranch)
	exception.ResponseIfNotFoundError(err)

	adminData, err := adminUsecase.AdminRepository.FindById(ctx, tx, updateRequestAdmin.Id, branchData.Name)
	exception.ResponseIfNotFoundError(err)

	addressData := adminUsecase.AddressUsecase.Update(ctx, updateRequestAdmin.Address)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(updateRequestAdmin.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	adminData.IdBranch = branchData.Id
	adminData.Code = updateRequestAdmin.Code
	adminData.Level = updateRequestAdmin.Level
	adminData.FullName = updateRequestAdmin.FullName
	adminData.Password = passwordHashedByte
	adminData.TelephoneNumber = updateRequestAdmin.TelephoneNumber
	adminData.Photo = updateRequestAdmin.Photo
	adminData.Status = updateRequestAdmin.Status
	adminData.UpdatedAt = time.Now()

	adminUsecase.AdminRepository.Update(ctx, tx, &adminData)
	return helper.ConvertToAdminResponse(&adminData, &addressData, &branchData.Name)
}

func (adminUsecase *AdminUsecaseImpl) Delete(ctx context.Context, adminId int) {
	tx, err := adminUsecase.DB.Begin()
	helper.PanicIfError(err)

	adminData, err := adminUsecase.AdminRepository.FindById(ctx, tx, adminId, "")
	exception.ResponseIfNotFoundError(err)

	adminUsecase.AdminRepository.Delete(ctx, tx, adminData.Id)
	helper.CommitOrRollback(tx)

	adminUsecase.AddressUsecase.Delete(ctx, adminData.IdAddress)
}
