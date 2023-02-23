package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/inventory"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type InventoryControllerImpl struct {
	InventoryUsecase usecase.InventoryUsecase
}

func NewInventoryControllerImpl(inventoryUsecase usecase.InventoryUsecase) *InventoryControllerImpl {
	return &InventoryControllerImpl{InventoryUsecase: inventoryUsecase}
}

func (inventoryController *InventoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allInventoryData := inventoryController.InventoryUsecase.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allInventoryData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (inventoryController *InventoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	inventoryIdInt, err := strconv.Atoi(p.ByName("inventoryId"))
	helper.PanicIfError(err)

	inventoryData := inventoryController.InventoryUsecase.FindById(r.Context(), inventoryIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    inventoryData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (inventoryController *InventoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createRequestInventory := inventory.CreateRequestInventory{}
	helper.ReadFromRequestBody(r, &createRequestInventory)

	inventoryData := inventoryController.InventoryUsecase.Create(r.Context(), &createRequestInventory)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    inventoryData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (inventoryController *InventoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var updateRequestInventory inventory.UpdateRequestInventory
	helper.ReadFromRequestBody(r, &updateRequestInventory)

	inventoryData := inventoryController.InventoryUsecase.Update(r.Context(), &updateRequestInventory)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    inventoryData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (inventoryController *InventoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	inventoryIdInt, err := strconv.Atoi(p.ByName("inventoryId"))
	helper.PanicIfError(err)
	inventoryController.InventoryUsecase.Delete(r.Context(), inventoryIdInt)

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)

}
