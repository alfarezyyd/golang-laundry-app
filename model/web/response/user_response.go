package response

import "time"

type UserResponse struct {
	Id              int              `json:"id,omitempty"`
	Level           string           `json:"level,omitempty"`
	FullName        string           `json:"full_name,omitempty"`
	Gender          string           `json:"gender,omitempty"`
	Password        string           `json:"password,omitempty"`
	Email           string           `json:"email,omitempty"`
	TelephoneNumber int              `json:"telephone_number,omitempty"`
	Photo           string           `json:"photo,omitempty"`
	EmailVerifiedAt *time.Time       `json:"email_verified_at,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	Address         *AddressResponse `json:"address,omitempty"`
}
