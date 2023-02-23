package inventory

type UpdateRequestInventory struct {
	Id         int    `json:"id,omitempty" validate:"required"`
	IdEmployee int    `json:"id_employee,omitempty" validate:"required"`
	Code       string `json:"code,omitempty" validate:"required"`
	Commodity  string `json:"commodity,omitempty" validate:"required"`
	Variant    string `json:"variant,omitempty" validate:"required"`
	Quantity   int    `json:"quantity,omitempty" validate:"required"`
	Price      int    `json:"price,omitempty" validate:"required"`
	Supplier   string `json:"supplier,omitempty" validate:"required"`
}
