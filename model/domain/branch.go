package domain

import "time"

type Branch struct {
	Id              int
	IdAddress       int
	Name            string
	TelephoneNumber int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
