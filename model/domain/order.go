package domain

import "time"

type Order struct {
	Id          int
	IdUser      int
	IdEmployee  int
	IdService   int
	Code        string
	Type        string
	Price       int
	Weight      int
	Payment     string
	Description string
	Status      string
	Entry       time.Time
	Out         time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
