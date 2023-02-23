package response

import "time"

type BranchResponse struct {
	Id              int              `json:"id,omitempty"`
	Code            string           `json:"code,omitempty"`
	Name            string           `json:"name,omitempty"`
	TelephoneNumber int              `json:"telephone_number,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	Address         *AddressResponse `json:"address,omitempty"`
}
