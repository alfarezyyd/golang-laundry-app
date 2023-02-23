package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/order"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"golang-laundry-app/usecase"
	"time"
)

type OrderUsecaseImpl struct {
	OrderRepository  repository.OrderRepository
	InventoryUsecase usecase.InventoryUsecase
	PromoUsecase     usecase.PromoUsecase
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrderUsecaseImpl(orderRepository repository.OrderRepository, inventoryUsecase usecase.InventoryUsecase, promoUsecase usecase.PromoUsecase, DB *sql.DB, validate *validator.Validate) *OrderUsecaseImpl {
	return &OrderUsecaseImpl{OrderRepository: orderRepository, InventoryUsecase: inventoryUsecase, PromoUsecase: promoUsecase, DB: DB, Validate: validate}
}

func (orderUsecase *OrderUsecaseImpl) FindAll(ctx context.Context) []response.OrderResponse {
	tx, err := orderUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	var allEmployeeName, allUserName, allServiceName []*string
	allOrderData := orderUsecase.OrderRepository.FindAll(ctx, tx, allEmployeeName, allUserName, allServiceName)

	var allOrderResponse []response.OrderResponse
	for index, orderData := range allOrderData {
		orderResponse := response.OrderResponse{
			Id:           orderData.Id,
			Code:         orderData.Code,
			Price:        orderData.Price,
			Status:       orderData.Status,
			UserName:     allUserName[index],
			EmployeeName: allEmployeeName[index],
			ServiceName:  allServiceName[index],
		}
		allOrderResponse = append(allOrderResponse, orderResponse)
	}
	return allOrderResponse
}

func (orderUsecase *OrderUsecaseImpl) FindById(ctx context.Context, orderId int) response.OrderResponse {
	tx, err := orderUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	var employeeName, userName, serviceName string
	orderData, err := orderUsecase.OrderRepository.FindById(ctx, tx, orderId, employeeName, userName, serviceName)
	exception.ResponseIfNotFoundError(err)

	allInventoryData := orderUsecase.InventoryUsecase.FindAllInventoryByOrder(ctx, &orderData.Id)
	allPromoData := orderUsecase.PromoUsecase.FindAllPromoByOrder(ctx, &orderData.Id)
	return helper.ConvertToOrderResponse(&orderData, allInventoryData, allPromoData, &userName, &employeeName, &serviceName)
}

func (orderUsecase *OrderUsecaseImpl) Create(ctx context.Context, orderCreateRequest *order.CreateRequestOrder) response.OrderResponse {
	err := orderUsecase.Validate.Struct(orderCreateRequest)
	helper.PanicIfError(err)

	tx, err := orderUsecase.DB.Begin()
	helper.PanicIfError(err)

	for _, idPromo := range orderCreateRequest.IdPromos {
		orderUsecase.PromoUsecase.FindById(ctx, idPromo)
	}

	for _, idInventory := range orderCreateRequest.IdInventories {
		orderUsecase.InventoryUsecase.FindById(ctx, idInventory)
	}

	entryDate, err := time.Parse("2006-01-02", orderCreateRequest.Entry)
	helper.PanicIfError(err)

	orderData := domain.Order{
		IdUser:      orderCreateRequest.IdUser,
		IdEmployee:  orderCreateRequest.IdEmployee,
		IdService:   orderCreateRequest.IdService,
		Code:        orderCreateRequest.Code,
		Type:        orderCreateRequest.Type,
		Price:       orderCreateRequest.Price,
		Weight:      orderCreateRequest.Weight,
		Payment:     orderCreateRequest.Payment,
		Description: orderCreateRequest.Description,
		Status:      orderCreateRequest.Status,
		Entry:       entryDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	orderUsecase.OrderRepository.Create(ctx, tx, &orderData)
	helper.CommitOrRollback(tx)

	newTx, err := orderUsecase.DB.Begin()
	defer helper.CommitOrRollback(newTx)
	helper.PanicIfError(err)

	for _, idPromo := range orderCreateRequest.IdPromos {
		orderUsecase.OrderRepository.CreateOrderPromo(ctx, newTx, &orderData.Id, &idPromo)
	}

	for _, idInventory := range orderCreateRequest.IdInventories {
		orderUsecase.OrderRepository.CreateOrderInventory(ctx, newTx, &orderData.Id, &idInventory)
	}
	return helper.ConvertToOrderResponse(&orderData, nil, nil, nil, nil, nil)
}

func (orderUsecase *OrderUsecaseImpl) Update(ctx context.Context, orderUpdateRequest *order.UpdateRequestOrder) response.OrderResponse {
	err := orderUsecase.Validate.Struct(orderUpdateRequest)
	helper.PanicIfError(err)

	tx, err := orderUsecase.DB.Begin()
	helper.PanicIfError(err)

	entryDate, err := time.Parse("2006-01-02", orderUpdateRequest.Entry)
	helper.PanicIfError(err)
	outDate, err := time.Parse("2006-01-02", orderUpdateRequest.Out)
	helper.PanicIfError(err)

	orderData, err := orderUsecase.OrderRepository.FindById(ctx, tx, orderUpdateRequest.Id, "", "", "")
	exception.ResponseIfNotFoundError(err)

	orderData.IdEmployee = orderUpdateRequest.IdEmployee
	orderData.IdService = orderUpdateRequest.IdService
	orderData.Type = orderUpdateRequest.Type
	orderData.Price = orderUpdateRequest.Price
	orderData.Weight = orderUpdateRequest.Weight
	orderData.Payment = orderUpdateRequest.Payment
	orderData.Description = orderUpdateRequest.Description
	orderData.Status = orderUpdateRequest.Status
	orderData.Entry = entryDate
	orderData.Out = outDate
	orderData.UpdatedAt = time.Now()

	orderUsecase.OrderRepository.Update(ctx, tx, &orderData)
	helper.CommitOrRollback(tx)

	newTx, err := orderUsecase.DB.Begin()
	defer helper.CommitOrRollback(newTx)
	helper.PanicIfError(err)

	if orderUpdateRequest.IdInventories != nil {
		orderUsecase.OrderRepository.DeleteOrderPromo(ctx, newTx, &orderData.Id)
		for _, idPromo := range orderUpdateRequest.IdPromos {
			orderUsecase.OrderRepository.CreateOrderPromo(ctx, newTx, &orderData.Id, &idPromo)
		}
	}

	if orderUpdateRequest.IdPromos != nil {
		orderUsecase.OrderRepository.DeleteOrderInventory(ctx, newTx, &orderData.Id)
		for _, idInventory := range orderUpdateRequest.IdInventories {
			orderUsecase.OrderRepository.CreateOrderInventory(ctx, newTx, &orderData.Id, &idInventory)
		}
	}

	allInventoryData := orderUsecase.InventoryUsecase.FindAllInventoryByOrder(ctx, &orderData.Id)
	allPromoData := orderUsecase.PromoUsecase.FindAllPromoByOrder(ctx, &orderData.Id)

	return helper.ConvertToOrderResponse(&orderData, allInventoryData, allPromoData, nil, nil, nil)
}

func (orderUsecase *OrderUsecaseImpl) Delete(ctx context.Context, orderId int) {
	tx, err := orderUsecase.DB.Begin()
	helper.PanicIfError(err)

	_, err = orderUsecase.OrderRepository.FindById(ctx, tx, orderId, "", "", "")
	exception.ResponseIfNotFoundError(err)

	orderUsecase.OrderRepository.Delete(ctx, tx, orderId)
	helper.CommitOrRollback(tx)

	newTx, err := orderUsecase.DB.Begin()
	helper.PanicIfError(err)

	orderUsecase.OrderRepository.DeleteOrderPromo(ctx, newTx, &orderId)
	orderUsecase.OrderRepository.DeleteOrderInventory(ctx, newTx, &orderId)
}

func (orderUsecase *OrderUsecaseImpl) DeleteAllOrderByUser(ctx context.Context, userId *int) {
	tx, err := orderUsecase.DB.Begin()
	helper.PanicIfError(err)

	allIdOrder := orderUsecase.OrderRepository.FindAllOrderByUser(ctx, tx, userId)

	for _, idOrder := range allIdOrder {
		orderUsecase.OrderRepository.DeleteOrderPromo(ctx, tx, &idOrder)
		orderUsecase.OrderRepository.DeleteOrderInventory(ctx, tx, &idOrder)
	}
	helper.CommitOrRollback(tx)

	newTx, err := orderUsecase.DB.Begin()
	defer helper.CommitOrRollback(newTx)
	helper.PanicIfError(err)

	orderUsecase.OrderRepository.DeleteAllOrderByUser(ctx, newTx, userId)
}
