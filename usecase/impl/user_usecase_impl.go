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
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/model/web/user"
	"golang-laundry-app/repository"
	"golang-laundry-app/usecase"
	"time"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	AddressUsecase usecase.AddressUsecase
	OrderUsecase   usecase.OrderUsecase
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserUsecaseImpl(userRepository repository.UserRepository, addressUsecase usecase.AddressUsecase, orderUsecase usecase.OrderUsecase, DB *sql.DB, validate *validator.Validate) *UserUsecaseImpl {
	return &UserUsecaseImpl{UserRepository: userRepository, AddressUsecase: addressUsecase, OrderUsecase: orderUsecase, DB: DB, Validate: validate}
}

func (userUsecase *UserUsecaseImpl) FindAll(ctx context.Context) []response.UserResponse {
	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allUserData := userUsecase.UserRepository.FindAll(ctx, tx)
	var allUserResponse []response.UserResponse
	for _, userData := range allUserData {
		userResponse := response.UserResponse{
			Id:       userData.Id,
			Level:    userData.Level,
			FullName: userData.FullName,
			Gender:   userData.Gender,
		}
		allUserResponse = append(allUserResponse, userResponse)
	}
	return allUserResponse
}

func (userUsecase *UserUsecaseImpl) FindById(ctx context.Context, userId int) response.UserResponse {
	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	userData, err := userUsecase.UserRepository.FindById(ctx, tx, userId)
	exception.ResponseIfNotFoundError(err)

	addressData := userUsecase.AddressUsecase.FindById(ctx, userData.IdAddress)
	return helper.ConvertToUserResponse(&userData, &addressData)
}

func (userUsecase *UserUsecaseImpl) Create(ctx context.Context, userCreateRequest *user.CreateRequestUser) response.UserResponse {
	err := userUsecase.Validate.Struct(userCreateRequest)
	helper.PanicIfError(err)

	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	addressData := userUsecase.AddressUsecase.Create(ctx, userCreateRequest.Address)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(userCreateRequest.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	userData := domain.User{
		IdAddress:       addressData.Id,
		Level:           userCreateRequest.Level,
		FullName:        userCreateRequest.FullName,
		Gender:          userCreateRequest.Gender,
		Password:        passwordHashedByte,
		Email:           userCreateRequest.Email,
		TelephoneNumber: userCreateRequest.TelephoneNumber,
		Photo:           userCreateRequest.Photo,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	userUsecase.UserRepository.Create(ctx, tx, &userData)
	return helper.ConvertToUserResponse(&userData, &addressData)
}

func (userUsecase *UserUsecaseImpl) Update(ctx context.Context, userUpdateRequest *user.UpdateRequestUser) response.UserResponse {
	err := userUsecase.Validate.Struct(userUpdateRequest)
	helper.PanicIfError(err)

	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	userData, err := userUsecase.UserRepository.FindById(ctx, tx, userUpdateRequest.Id)
	exception.ResponseIfNotFoundError(err)

	addressData := userUsecase.AddressUsecase.Update(ctx, userUpdateRequest.Address)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(userUpdateRequest.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	userData.Level = userUpdateRequest.Level
	userData.FullName = userUpdateRequest.FullName
	userData.Gender = userUpdateRequest.Gender
	userData.Password = passwordHashedByte
	userData.TelephoneNumber = userUpdateRequest.TelephoneNumber
	userData.Photo = userUpdateRequest.Photo
	userData.UpdatedAt = time.Now()

	userUsecase.UserRepository.Update(ctx, tx, &userData)
	return helper.ConvertToUserResponse(&userData, &addressData)

}

func (userUsecase *UserUsecaseImpl) Delete(ctx context.Context, userId int) {
	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)

	_, err = userUsecase.UserRepository.FindById(ctx, tx, userId)
	exception.ResponseIfNotFoundError(err)

	userUsecase.OrderUsecase.DeleteAllOrderByUser(ctx, &userId)
	userUsecase.UserRepository.Delete(ctx, tx, userId)
}

func (userUsecase *UserUsecaseImpl) Login(ctx context.Context, userLoginRequest *user.LoginRequestUser) response.UserResponse {
	err := userUsecase.Validate.Struct(userLoginRequest)
	helper.PanicIfError(err)

	tx, err := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	passwordHashing := sha256.New()
	passwordHashing.Write([]byte(userLoginRequest.Password))
	passwordHashed := passwordHashing.Sum(nil)
	passwordHashedByte := fmt.Sprintf("%x", passwordHashed)

	userData, err := userUsecase.UserRepository.FindAuthUser(ctx, tx, userLoginRequest.Email, passwordHashedByte)
	exception.ResponseIfNotFoundError(err)

	return response.UserResponse{
		Email: userData.Email,
	}
}
