package employee

import "golang-laundry-app/model/web/address"

type CreateRequestEmployee struct {
	IdBranch        int                           `json:"id_branch,omitempty" validate:"required"`
	Code            string                        `json:"code,omitempty" validate:"required"`
	Level           string                        `json:"level,omitempty" validate:"required"`
	FullName        string                        `json:"full_name,omitempty" validate:"required"`
	Password        string                        `json:"password,omitempty" validate:"required"`
	Email           string                        `json:"email,omitempty" validate:"required,email"`
	TelephoneNumber int                           `json:"telephone_number,omitempty" validate:"required`
	Photo           string                        `json:"photo,omitempty" validate:"required"`
	Status          string                        `json:"status,omitempty" validate:"required"`
	Address         *address.CreateRequestAddress `json:"address" validate:"required"`
}
