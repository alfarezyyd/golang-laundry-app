package service

type UpdateRequestService struct {
	Id       int    `json:"id,omitempty" validate:"required"`
	Code     string `json:"code,omitempty" validate:"required"`
	Name     string `json:"name,omitempty" validate:"required"`
	Price    int    `json:"price,omitempty" validate:"required"`
	Duration int    `json:"duration,omitempty" validate:"required"`
}
