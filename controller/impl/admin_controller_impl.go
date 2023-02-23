package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/admin"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type AdminControllerImpl struct {
	AdminUsecase usecase.AdminUsecase
}

func NewAdminControllerImpl(adminUsecase usecase.AdminUsecase) *AdminControllerImpl {
	return &AdminControllerImpl{AdminUsecase: adminUsecase}
}

func (adminController *AdminControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allAdminData := adminController.AdminUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allAdminData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (adminController *AdminControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	adminIdInt, err := strconv.Atoi(p.ByName("adminId"))
	helper.PanicIfError(err)

	adminData := adminController.AdminUsecase.FindById(r.Context(), adminIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    adminData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (adminController *AdminControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	adminCreateRequest := admin.CreateRequestAdmin{}
	helper.ReadFromRequestBody(r, &adminCreateRequest)

	adminData := adminController.AdminUsecase.Create(r.Context(), &adminCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    adminData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (adminController *AdminControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	adminUpdateRequest := admin.UpdateRequestAdmin{}
	helper.ReadFromRequestBody(r, &adminUpdateRequest)

	adminData := adminController.AdminUsecase.Update(r.Context(), &adminUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    adminData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (adminController *AdminControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	adminId, err := strconv.Atoi(p.ByName("adminId"))
	helper.PanicIfError(err)

	adminController.AdminUsecase.Delete(r.Context(), adminId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
