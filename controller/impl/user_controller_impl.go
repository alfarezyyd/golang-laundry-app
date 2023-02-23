package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/user"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserControllerImpl(userUsecase usecase.UserUsecase) *UserControllerImpl {
	return &UserControllerImpl{UserUsecase: userUsecase}
}

func (userController *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allUserData := userController.UserUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allUserData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userIdInt, err := strconv.Atoi(p.ByName("userId"))
	helper.PanicIfError(err)

	userData := userController.UserUsecase.FindById(r.Context(), userIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    userData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateRequest := user.CreateRequestUser{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	userData := userController.UserUsecase.Create(r.Context(), &userCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    userData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdateRequest := user.UpdateRequestUser{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userData := userController.UserUsecase.Update(r.Context(), &userUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    userData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := strconv.Atoi(p.ByName("userId"))
	helper.PanicIfError(err)

	userController.UserUsecase.Delete(r.Context(), userId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
