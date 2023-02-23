package promo

type CreateRequestPromo struct {
	Code        string `json:"code,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Discount    int    `json:"discount,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Status      string `json:"status,omitempty" validate:"required"`
	Photo       string `json:"photo,omitempty" validate:"required"`
	Start       string `json:"start" validate:"required"`
	End         string `json:"end" validate:"required"`
}
