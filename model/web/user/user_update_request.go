package user

import "golang-laundry-app/model/web/address"

type UpdateRequestUser struct {
	Id              int                           `json:"id,omitempty" validate:"required"`
	Level           string                        `json:"level,omitempty"`
	FullName        string                        `json:"full_name,omitempty"`
	Gender          string                        `json:"gender,omitempty"`
	Password        string                        `json:"password,omitempty"`
	Email           string                        `json:"email,omitempty"`
	TelephoneNumber int                           `json:"telephone_number,omitempty"`
	Photo           string                        `json:"photo,omitempty"`
	Address         *address.UpdateRequestAddress `json:"address,omitempty"`
}
