package helper

import (
	"errors"
	"time"
)

func DateCompare(startDate, endDate time.Time) error {
	if startDate.After(endDate) {
		return errors.New("date not valid")
	}
	return nil
}
