package branch

import (
	"golang-laundry-app/model/web/address"
)

type UpdateRequestBranch struct {
	Id              int                           `json:"id,omitempty" validate:"required"`
	Name            string                        `json:"name,omitempty" validate:"required,min=10"`
	TelephoneNumber int                           `json:"telephone_number,omitempty" validate:"required"`
	Address         *address.UpdateRequestAddress `json:"address,omitempty" validate:"required"`
}
