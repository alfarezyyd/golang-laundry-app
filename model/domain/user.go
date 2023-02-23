package domain

import "time"

type User struct {
	Id              int
	IdAddress       int
	Level           string
	FullName        string
	Gender          string
	Password        string
	Email           string
	TelephoneNumber int
	Photo           string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
