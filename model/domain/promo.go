package domain

import "time"

type Promo struct {
	Id          int
	Code        string
	Name        string
	Discount    int
	Description string
	Status      string
	Photo       string
	Start       time.Time
	End         time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
