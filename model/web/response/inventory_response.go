package response

import "time"

type InventoryResponse struct {
	Id        int               `json:"id,omitempty"`
	Code      string            `json:"code,omitempty"`
	Commodity string            `json:"commodity,omitempty"`
	Variant   string            `json:"variant,omitempty"`
	Quantity  int               `json:"quantity,omitempty"`
	Price     int               `json:"price,omitempty"`
	Supplier  string            `json:"supplier,omitempty"`
	CreatedAt *time.Time        `json:"created_at,omitempty"`
	UpdatedAt *time.Time        `json:"updated_at,omitempty"`
	Employee  *EmployeeResponse `json:"employee,omitempty"`
}
