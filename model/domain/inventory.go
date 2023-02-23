package domain

import "time"

type Inventory struct {
	Id         int
	IdEmployee int
	Code       string
	Commodity  string
	Variant    string
	Quantity   int
	Price      int
	Supplier   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
