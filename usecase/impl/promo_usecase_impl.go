package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang-laundry-app/exception"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/promo"
	"golang-laundry-app/model/web/response"
	"golang-laundry-app/repository"
	"time"
)

type PromoUsecaseImpl struct {
	PromoRepository repository.PromoRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewPromoUsecaseImpl(promoRepository repository.PromoRepository, DB *sql.DB, validate *validator.Validate) *PromoUsecaseImpl {
	return &PromoUsecaseImpl{PromoRepository: promoRepository, DB: DB, Validate: validate}
}

func (promoUsecase *PromoUsecaseImpl) FindAll(ctx context.Context) []response.PromoResponse {
	tx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allPromoData := promoUsecase.PromoRepository.FindAll(ctx, tx)
	var allPromoResponse []response.PromoResponse
	for _, promoData := range allPromoData {
		promoResponse := response.PromoResponse{
			Id:       promoData.Id,
			Code:     promoData.Code,
			Name:     promoData.Name,
			Discount: promoData.Discount,
		}
		allPromoResponse = append(allPromoResponse, promoResponse)
	}
	return allPromoResponse
}

func (promoUsecase *PromoUsecaseImpl) FindById(ctx context.Context, promoId int) response.PromoResponse {
	tx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	promoData, err := promoUsecase.PromoRepository.FindById(ctx, tx, promoId)
	exception.ResponseIfNotFoundError(err)

	return helper.ConvertToPromoResponse(&promoData)
}

func (promoUsecase *PromoUsecaseImpl) Create(ctx context.Context, promoCreateRequest *promo.CreateRequestPromo) response.PromoResponse {
	err := promoUsecase.Validate.Struct(promoCreateRequest)
	helper.PanicIfError(err)

	tx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	startDate, err := time.Parse("2006-01-02", promoCreateRequest.Start)
	helper.PanicIfError(err)
	endDate, err := time.Parse("2006-01-02", promoCreateRequest.End)
	helper.PanicIfError(err)

	err = helper.DateCompare(startDate, endDate)
	exception.ResponseIfProgramError(err)

	promoData := domain.Promo{
		Code:        promoCreateRequest.Code,
		Name:        promoCreateRequest.Name,
		Discount:    promoCreateRequest.Discount,
		Description: promoCreateRequest.Description,
		Status:      promoCreateRequest.Status,
		Photo:       promoCreateRequest.Photo,
		Start:       startDate,
		End:         endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	promoUsecase.PromoRepository.Create(ctx, tx, &promoData)
	return helper.ConvertToPromoResponse(&promoData)
}

func (promoUsecase *PromoUsecaseImpl) Update(ctx context.Context, promoUpdateRequest *promo.UpdateRequestPromo) response.PromoResponse {
	err := promoUsecase.Validate.Struct(promoUpdateRequest)
	helper.PanicIfError(err)

	tx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	startDate, err := time.Parse("2006-01-02", promoUpdateRequest.Start)
	helper.PanicIfError(err)
	endDate, err := time.Parse("2006-01-02", promoUpdateRequest.End)
	helper.PanicIfError(err)

	err = helper.DateCompare(startDate, endDate)
	exception.ResponseIfProgramError(err)

	promoData := domain.Promo{
		Code:        promoUpdateRequest.Code,
		Name:        promoUpdateRequest.Name,
		Discount:    promoUpdateRequest.Discount,
		Description: promoUpdateRequest.Description,
		Status:      promoUpdateRequest.Status,
		Photo:       promoUpdateRequest.Photo,
		Start:       startDate,
		End:         endDate,
	}
	promoUsecase.PromoRepository.Update(ctx, tx, &promoData)
	return helper.ConvertToPromoResponse(&promoData)

}

func (promoUsecase *PromoUsecaseImpl) Delete(ctx context.Context, promoId int) {
	tx, err := promoUsecase.DB.Begin()
	helper.PanicIfError(err)

	promoUsecase.FindById(ctx, promoId)
	promoUsecase.PromoRepository.DeletePromoOrder(ctx, tx, &promoId)
	helper.CommitOrRollback(tx)

	newTx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(newTx)
	helper.PanicIfError(err)

	promoUsecase.PromoRepository.Delete(ctx, newTx, promoId)
}

func (promoUsecase *PromoUsecaseImpl) FindAllPromoByOrder(ctx context.Context, orderId *int) []*response.PromoResponse {
	tx, err := promoUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	allPromoData := promoUsecase.PromoRepository.FindAllPromoByOrder(ctx, tx, orderId)
	var allPromoResponse []*response.PromoResponse
	for _, promoData := range allPromoData {
		promoResponse := response.PromoResponse{
			Id:       promoData.Id,
			Code:     promoData.Code,
			Name:     promoData.Name,
			Discount: promoData.Discount,
		}
		allPromoResponse = append(allPromoResponse, &promoResponse)
	}
	return allPromoResponse

}
