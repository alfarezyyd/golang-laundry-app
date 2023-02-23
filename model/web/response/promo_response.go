package response

import "time"

type PromoResponse struct {
	Id          int        `json:"id,omitempty"`
	Code        string     `json:"code,omitempty"`
	Name        string     `json:"name,omitempty"`
	Discount    int        `json:"discount,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	Photo       string     `json:"photo,omitempty"`
	Start       *time.Time `json:"start,omitempty"`
	End         *time.Time `json:"end,omitempty"`
}
