package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/address"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
)

type AddressUsecaseImpl struct {
	AddressRepository repository.AddressRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAddressUsecaseImpl(addressRepository repository.AddressRepository, DB *sql.DB, validate *validator.Validate) *AddressUsecaseImpl {
	return &AddressUsecaseImpl{AddressRepository: addressRepository, DB: DB, Validate: validate}
}

func (addressUsecase *AddressUsecaseImpl) FindById(ctx context.Context, addressId int) response.AddressResponse {
	tx, err := addressUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressData, err := addressUsecase.AddressRepository.FindById(ctx, tx, addressId)
	exception.ResponseIfNotFoundError(err)

	return helper.ConvertToAddressResponse(&addressData)
}

func (addressUsecase *AddressUsecaseImpl) Create(ctx context.Context, addressCreateRequest *address.CreateRequestAddress) response.AddressResponse {
	err := addressUsecase.Validate.Struct(addressCreateRequest)
	helper.PanicIfError(err)

	tx, err := addressUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressData := domain.Address{
		BuildingName:        addressCreateRequest.BuildingName,
		BuildingNumber:      addressCreateRequest.BuildingNumber,
		Street:              addressCreateRequest.Street,
		Village:             addressCreateRequest.Village,
		NeighbourhoodNumber: addressCreateRequest.NeighbourhoodNumber,
		HamletNumber:        addressCreateRequest.HamletNumber,
		SubDistrict:         addressCreateRequest.SubDistrict,
		District:            addressCreateRequest.District,
		Province:            addressCreateRequest.Province,
		PostalCode:          addressCreateRequest.PostalCode,
		Description:         addressCreateRequest.Description,
	}

	addressUsecase.AddressRepository.Create(ctx, tx, &addressData)
	return helper.ConvertToAddressResponse(&addressData)
}

func (addressUsecase *AddressUsecaseImpl) Update(ctx context.Context, addressUpdateRequest *address.UpdateRequestAddress) response.AddressResponse {
	err := addressUsecase.Validate.Struct(addressUpdateRequest)
	helper.PanicIfError(err)

	tx, err := addressUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressUsecase.FindById(ctx, addressUpdateRequest.Id)

	addressData := domain.Address{
		Id:                  addressUpdateRequest.Id,
		BuildingName:        addressUpdateRequest.BuildingName,
		BuildingNumber:      addressUpdateRequest.BuildingNumber,
		Street:              addressUpdateRequest.Street,
		Village:             addressUpdateRequest.Village,
		NeighbourhoodNumber: addressUpdateRequest.NeighbourhoodNumber,
		HamletNumber:        addressUpdateRequest.HamletNumber,
		SubDistrict:         addressUpdateRequest.SubDistrict,
		District:            addressUpdateRequest.District,
		Province:            addressUpdateRequest.Province,
		PostalCode:          addressUpdateRequest.PostalCode,
		Description:         addressUpdateRequest.Description,
	}

	addressUsecase.AddressRepository.Create(ctx, tx, &addressData)
	return helper.ConvertToAddressResponse(&addressData)
}

func (addressUsecase *AddressUsecaseImpl) Delete(ctx context.Context, addressId int) {
	tx, err := addressUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressUsecase.FindById(ctx, addressId)
	addressUsecase.AddressRepository.Delete(ctx, tx, addressId)
}
