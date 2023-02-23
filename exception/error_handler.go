package exception

import (
	"github.com/go-playground/validator"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if validationErrors(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	if programError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exceptionType, state := err.(validator.ValidationErrors)
	if state {
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    exceptionType.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exceptionType, state := err.(NotFoundError)
	if state {
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Message: "Not Found Error",
			Data:    exceptionType.Error,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func programError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exceptionType, state := err.(ProgramError)
	if state {
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: exceptionType.Error,
			Data:    exceptionType.Error,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    err,
	}
	helper.WriteToResponseBody(w, webResponse)
}
