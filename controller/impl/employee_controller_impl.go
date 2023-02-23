package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/employee"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type EmployeeControllerImpl struct {
	EmployeeUsecase usecase.EmployeeUsecase
}

func NewEmployeeControllerImpl(employeeUsecase usecase.EmployeeUsecase) *EmployeeControllerImpl {
	return &EmployeeControllerImpl{EmployeeUsecase: employeeUsecase}
}

func (employeeController *EmployeeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allEmployeeData := employeeController.EmployeeUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allEmployeeData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (employeeController *EmployeeControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeIdInt, err := strconv.Atoi(p.ByName("employeeId"))
	helper.PanicIfError(err)

	employeeData := employeeController.EmployeeUsecase.FindById(r.Context(), employeeIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    employeeData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (employeeController *EmployeeControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var createRequestEmployee employee.CreateRequestEmployee
	helper.ReadFromRequestBody(r, &createRequestEmployee)

	employeeData := employeeController.EmployeeUsecase.Create(r.Context(), &createRequestEmployee)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    employeeData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (employeeController *EmployeeControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	updateRequestEmployee := employee.UpdateRequestEmployee{}
	helper.ReadFromRequestBody(r, &updateRequestEmployee)

	employeeData := employeeController.EmployeeUsecase.Update(r.Context(), &updateRequestEmployee)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    employeeData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (employeeController *EmployeeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeIdInt, err := strconv.Atoi(p.ByName("employeeId"))
	helper.PanicIfError(err)
	employeeController.EmployeeUsecase.Delete(r.Context(), employeeIdInt)

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)

}
