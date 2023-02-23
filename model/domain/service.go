package domain

import "time"

type Service struct {
	Id        int
	Code      string
	Name      string
	Price     int
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
