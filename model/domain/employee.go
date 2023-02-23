package domain

import "time"

type Employee struct {
	Id              int
	IdAddress       int
	IdBranch        int
	Code            string
	Level           string
	FullName        string
	Password        string
	Email           string
	TelephoneNumber int
	Photo           string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
