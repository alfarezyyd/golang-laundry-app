package response

import (
	"time"
)

type OrderResponse struct {
	Id                int                  `json:"id,omitempty"`
	UserName          *string              `json:"user_name,omitempty"`
	EmployeeName      *string              `json:"employee_name,omitempty"`
	ServiceName       *string              `json:"service_name,omitempty"`
	Code              string               `json:"code,omitempty"`
	Type              string               `json:"type,omitempty"`
	Price             int                  `json:"price,omitempty"`
	Weight            int                  `json:"weight,omitempty"`
	Payment           string               `json:"payment,omitempty"`
	Description       string               `json:"description,omitempty"`
	Status            string               `json:"status,omitempty"`
	InventoryResponse []*InventoryResponse `json:"inventory_response,omitempty"`
	PromoResponse     []*PromoResponse     `json:"promo_response,omitempty"`
	Entry             *time.Time           `json:"entry,omitempty"`
	Out               *time.Time           `json:"out,omitempty"`
	CreatedAt         *time.Time           `json:"created_at,omitempty"`
	UpdatedAt         *time.Time           `json:"updated_at,omitempty"`
}
