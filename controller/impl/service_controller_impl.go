package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/service"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type ServiceControllerImpl struct {
	ServiceUsecase usecase.ServiceUsecase
}

func NewServiceControllerImpl(serviceUsecase usecase.ServiceUsecase) *ServiceControllerImpl {
	return &ServiceControllerImpl{ServiceUsecase: serviceUsecase}
}

func (serviceController *ServiceControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allServiceData := serviceController.ServiceUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allServiceData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (serviceController *ServiceControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serviceIdInt, err := strconv.Atoi(p.ByName("serviceId"))
	helper.PanicIfError(err)

	serviceData := serviceController.ServiceUsecase.FindById(r.Context(), serviceIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    serviceData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (serviceController *ServiceControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serviceCreateRequest := service.CreateRequestService{}
	helper.ReadFromRequestBody(r, &serviceCreateRequest)

	serviceData := serviceController.ServiceUsecase.Create(r.Context(), &serviceCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    serviceData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (serviceController *ServiceControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serviceUpdateRequest := service.UpdateRequestService{}
	helper.ReadFromRequestBody(r, &serviceUpdateRequest)

	serviceData := serviceController.ServiceUsecase.Update(r.Context(), &serviceUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    serviceData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (serviceController *ServiceControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serviceId, err := strconv.Atoi(p.ByName("serviceId"))
	helper.PanicIfError(err)

	serviceController.ServiceUsecase.Delete(r.Context(), serviceId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
