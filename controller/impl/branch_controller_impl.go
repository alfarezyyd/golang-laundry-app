package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/branch"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type BranchControllerImpl struct {
	BranchUsecase usecase.BranchUsecase
}

func NewBranchControllerImpl(branchUsecase usecase.BranchUsecase) *BranchControllerImpl {
	return &BranchControllerImpl{BranchUsecase: branchUsecase}
}

func (branchController *BranchControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allBranchData := branchController.BranchUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allBranchData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (branchController *BranchControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branchIdInt, err := strconv.Atoi(p.ByName("branchId"))
	helper.PanicIfError(err)

	branchData := branchController.BranchUsecase.FindById(r.Context(), branchIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    branchData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (branchController *BranchControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branchCreateRequest := branch.CreateRequestBranch{}
	helper.ReadFromRequestBody(r, &branchCreateRequest)

	branchData := branchController.BranchUsecase.Create(r.Context(), &branchCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    branchData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (branchController *BranchControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branchUpdateRequest := branch.UpdateRequestBranch{}
	helper.ReadFromRequestBody(r, &branchUpdateRequest)

	branchData := branchController.BranchUsecase.Update(r.Context(), &branchUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    branchData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (branchController *BranchControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branchId, err := strconv.Atoi(p.ByName("branchId"))
	helper.PanicIfError(err)

	branchController.BranchUsecase.Delete(r.Context(), branchId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
