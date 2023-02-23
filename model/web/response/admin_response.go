package response

import "time"

type AdminResponse struct {
	Id              int              `json:"id,omitempty"`
	Code            string           `json:"code,omitempty"`
	Level           string           `json:"level,omitempty"`
	FullName        string           `json:"full_name,omitempty"`
	Password        string           `json:"password,omitempty"`
	Email           string           `json:"email,omitempty"`
	TelephoneNumber int              `json:"telephone_number,omitempty"`
	Photo           string           `json:"photo,omitempty"`
	Status          string           `json:"status,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	Address         *AddressResponse `json:"address,omitempty"`
	BranchName      string           `json:"branch_name,omitempty"`
}
