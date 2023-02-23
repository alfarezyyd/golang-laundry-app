package app

import (
	"database/sql"
	"golang-laundry-app/helper"
	"time"
)

func NewSetupDatabase() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_laundry?parseTime=true")
	helper.PanicIfError(err)

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
