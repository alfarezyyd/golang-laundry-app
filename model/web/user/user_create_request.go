package user

import "golang-laundry-app/model/web/address"

type CreateRequestUser struct {
	Level           string                        `json:"level,omitempty" validate:"required"`
	FullName        string                        `json:"full_name,omitempty" validate:"required"`
	Gender          string                        `json:"gender,omitempty" validate:"required"`
	Password        string                        `json:"password,omitempty" validate:"required"`
	Email           string                        `json:"email,omitempty" validate:"required"`
	TelephoneNumber int                           `json:"telephone_number,omitempty" validate:"required"`
	Photo           string                        `json:"photo,omitempty" validate:"required"`
	Address         *address.CreateRequestAddress `json:"address" validate:"required"`
}
