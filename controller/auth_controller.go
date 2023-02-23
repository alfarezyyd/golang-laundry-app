package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
