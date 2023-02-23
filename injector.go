//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/app"
	"golang-laundry-app/controller"
	controllerImpl "golang-laundry-app/controller/impl"
	"golang-laundry-app/repository"
	repositoryImpl "golang-laundry-app/repository/impl"
	"golang-laundry-app/usecase"
	usecaseImpl "golang-laundry-app/usecase/impl"
)

var addressSet = wire.NewSet(
	repositoryImpl.NewAddressRepositoryImpl,
	wire.Bind(new(repository.AddressRepository), new(*repositoryImpl.AddressRepositoryImpl)),
	usecaseImpl.NewAddressUsecaseImpl,
	wire.Bind(new(usecase.AddressUsecase), new(*usecaseImpl.AddressUsecaseImpl)),
)

var employeeSet = wire.NewSet(
	repositoryImpl.NewEmployeeRepositoryImpl,
	wire.Bind(new(repository.EmployeeRepository), new(*repositoryImpl.EmployeeRepositoryImpl)),
	usecaseImpl.NewEmployeeUsecaseImpl,
	wire.Bind(new(usecase.EmployeeUsecase), new(*usecaseImpl.EmployeeUsecaseImpl)),
	controllerImpl.NewEmployeeControllerImpl,
	wire.Bind(new(app.EmployeeController), new(*controllerImpl.EmployeeControllerImpl)),
)

var branchSet = wire.NewSet(
	repositoryImpl.NewBranchRepositoryImpl,
	wire.Bind(new(repository.BranchRepository), new(*repositoryImpl.BranchRepositoryImpl)),
	usecaseImpl.NewBranchUsecaseImpl,
	wire.Bind(new(usecase.BranchUsecase), new(*usecaseImpl.BranchUsecaseImpl)),
	controllerImpl.NewBranchControllerImpl,
	wire.Bind(new(app.BranchController), new(*controllerImpl.BranchControllerImpl)),
)

var inventorySet = wire.NewSet(
	repositoryImpl.NewInventoryRepositoryImpl,
	wire.Bind(new(repository.InventoryRepository), new(*repositoryImpl.InventoryRepositoryImpl)),
	usecaseImpl.NewInventoryUsecaseImpl,
	wire.Bind(new(usecase.InventoryUsecase), new(*usecaseImpl.InventoryUsecaseImpl)),
	controllerImpl.NewInventoryControllerImpl,
	wire.Bind(new(app.InventoryController), new(*controllerImpl.InventoryControllerImpl)),
)

var serviceSet = wire.NewSet(
	repositoryImpl.NewServiceRepositoryImpl,
	wire.Bind(new(repository.ServiceRepository), new(*repositoryImpl.ServiceRepositoryImpl)),
	usecaseImpl.NewServiceUsecaseImpl,
	wire.Bind(new(usecase.ServiceUsecase), new(*usecaseImpl.ServiceUsecaseImpl)),
	controllerImpl.NewServiceControllerImpl,
	wire.Bind(new(app.ServiceController), new(*controllerImpl.ServiceControllerImpl)),
)

var promoSet = wire.NewSet(
	repositoryImpl.NewPromoRepositoryImpl,
	wire.Bind(new(repository.PromoRepository), new(*repositoryImpl.PromoRepositoryImpl)),
	usecaseImpl.NewPromoUsecaseImpl,
	wire.Bind(new(usecase.PromoUsecase), new(*usecaseImpl.PromoUsecaseImpl)),
	controllerImpl.NewPromoControllerImpl,
	wire.Bind(new(app.PromoController), new(*controllerImpl.PromoControllerImpl)),
)

var userSet = wire.NewSet(
	repositoryImpl.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repositoryImpl.UserRepositoryImpl)),
	usecaseImpl.NewUserUsecaseImpl,
	wire.Bind(new(usecase.UserUsecase), new(*usecaseImpl.UserUsecaseImpl)),
	controllerImpl.NewUserControllerImpl,
	wire.Bind(new(app.UserController), new(*controllerImpl.UserControllerImpl)),
)

var adminSet = wire.NewSet(
	repositoryImpl.NewAdminRepositoryImpl,
	wire.Bind(new(repository.AdminRepository), new(*repositoryImpl.AdminRepositoryImpl)),
	usecaseImpl.NewAdminUsecaseImpl,
	wire.Bind(new(usecase.AdminUsecase), new(*usecaseImpl.AdminUsecaseImpl)),
	controllerImpl.NewAdminControllerImpl,
	wire.Bind(new(app.AdminController), new(*controllerImpl.AdminControllerImpl)),
)

var orderSet = wire.NewSet(
	repositoryImpl.NewOrderRepositoryImpl,
	wire.Bind(new(repository.OrderRepository), new(*repositoryImpl.OrderRepositoryImpl)),
	usecaseImpl.NewOrderUsecaseImpl,
	wire.Bind(new(usecase.OrderUsecase), new(*usecaseImpl.OrderUsecaseImpl)),
	controllerImpl.NewOrderControllerImpl,
	wire.Bind(new(app.OrderController), new(*controllerImpl.OrderControllerImpl)),
)

var authSet = wire.NewSet(
	controllerImpl.NewAuthControllerImpl,
	wire.Bind(new(controller.AuthController), new(*controllerImpl.AuthControllerImpl)),
)

var allSet = wire.NewSet(
	addressSet, employeeSet, branchSet, inventorySet, serviceSet, promoSet, adminSet, userSet, orderSet, authSet,
)

func InitializedHandler() *httprouter.Router {
	wire.Build(app.NewSetupDatabase, app.NewSetupRouter, validator.New, allSet)
	return nil
}
