package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-laundry-app/helper"
	"golang-laundry-app/middleware"
	"net/http"
)

func main() {
	router := InitializedHandler()
	newRouter := middleware.NewGatewayMiddleware(router)
	newRouter.RegisterMiddleware(middleware.AuthMiddleware)
	server := http.Server{
		Addr:    ":3000",
		Handler: newRouter,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
