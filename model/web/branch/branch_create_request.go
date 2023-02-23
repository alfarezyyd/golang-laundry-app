package branch

import "golang-laundry-app/model/web/address"

type CreateRequestBranch struct {
	Name            string                        `json:"name,omitempty" validate:"required,min=10"`
	TelephoneNumber int                           `json:"telephone_number,omitempty" validate:"required"`
	Address         *address.CreateRequestAddress `json:"address,omitempty" validate:"required"`
}
