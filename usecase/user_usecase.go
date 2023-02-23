package usecase

import (
	"context"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/model/web/user"
)

type UserUsecase interface {
	FindAll(ctx context.Context) []response.UserResponse
	FindById(ctx context.Context, userId int) response.UserResponse
	Create(ctx context.Context, userCreateRequest *user.CreateRequestUser) response.UserResponse
	Update(ctx context.Context, userUpdateRequest *user.UpdateRequestUser) response.UserResponse
	Delete(ctx context.Context, userId int)
	Login(ctx context.Context, userLoginRequest *user.LoginRequestUser) response.UserResponse
}
