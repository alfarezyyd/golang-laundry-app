package impl

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/promo"
	"golang-laundry-app/usecase"
	"net/http"
	"strconv"
)

type PromoControllerImpl struct {
	PromoUsecase usecase.PromoUsecase
}

func NewPromoControllerImpl(promoUsecase usecase.PromoUsecase) *PromoControllerImpl {
	return &PromoControllerImpl{PromoUsecase: promoUsecase}
}

func (promoController *PromoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	allPromoData := promoController.PromoUsecase.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    allPromoData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (promoController *PromoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	promoIdInt, err := strconv.Atoi(p.ByName("promoId"))
	helper.PanicIfError(err)

	promoData := promoController.PromoUsecase.FindById(r.Context(), promoIdInt)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    promoData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (promoController *PromoControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	promoCreateRequest := promo.CreateRequestPromo{}
	helper.ReadFromRequestBody(r, &promoCreateRequest)

	promoData := promoController.PromoUsecase.Create(r.Context(), &promoCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    promoData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (promoController *PromoControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	promoUpdateRequest := promo.UpdateRequestPromo{}
	helper.ReadFromRequestBody(r, &promoUpdateRequest)

	promoData := promoController.PromoUsecase.Update(r.Context(), &promoUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    promoData,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (promoController *PromoControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	promoId, err := strconv.Atoi(p.ByName("promoId"))
	helper.PanicIfError(err)

	promoController.PromoUsecase.Delete(r.Context(), promoId)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success!",
		Data:    nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}
