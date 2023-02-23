package response

import "time"

type ServiceResponse struct {
	Id        int       `json:"id,omitempty"`
	Code      string    `json:"code,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     int       `json:"price,omitempty"`
	Duration  int       `json:"duration,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
