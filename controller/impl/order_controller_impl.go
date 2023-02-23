package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/order"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type OrderControllerImpl struct {
	OrderUsecase usecase.OrderUsecase
}

func NewOrderControllerImpl(orderUsecase usecase.OrderUsecase) *OrderControllerImpl {
	return &OrderControllerImpl{OrderUsecase: orderUsecase}
}

func (orderController *OrderControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allOrderData := orderController.OrderUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allOrderData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (orderController *OrderControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderIdInt, err := strconv.Atoi(p.ByName("orderId"))
	helper.PanicIfError(err)

	orderData := orderController.OrderUsecase.FindById(r.Context(), orderIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    orderData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (orderController *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderCreateRequest := order.CreateRequestOrder{}
	helper.ReadFromRequestBody(r, &orderCreateRequest)

	orderData := orderController.OrderUsecase.Create(r.Context(), &orderCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    orderData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (orderController *OrderControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderUpdateRequest := order.UpdateRequestOrder{}
	helper.ReadFromRequestBody(r, &orderUpdateRequest)

	orderData := orderController.OrderUsecase.Update(r.Context(), &orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    orderData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (orderController *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderId, err := strconv.Atoi(p.ByName("orderId"))
	helper.PanicIfError(err)

	orderController.OrderUsecase.Delete(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
