package admin

import (
	"golang-laundry-app/model/web/address"
)

type UpdateRequestAdmin struct {
	Id              int                           `json:"id,omitempty"`
	IdBranch        int                           `json:"id_branch,omitempty"`
	Code            string                        `json:"code,omitempty"`
	Level           string                        `json:"level,omitempty"`
	FullName        string                        `json:"full_name,omitempty"`
	Password        string                        `json:"password,omitempty"`
	Email           string                        `json:"email,omitempty"`
	TelephoneNumber int                           `json:"telephone_number,omitempty"`
	Photo           string                        `json:"photo,omitempty"`
	Status          string                        `json:"status,omitempty"`
	Address         *address.UpdateRequestAddress `json:"address,omitempty"`
}
