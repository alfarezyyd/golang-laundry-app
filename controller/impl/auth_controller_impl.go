package impl

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/config"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"golang-laundry-app/model/web/user"
	"golang-laundry-app/usecase"
	"net/http"
	"time"
)

type AuthControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthControllerImpl(userUsecase usecase.UserUsecase) *AuthControllerImpl {
	return &AuthControllerImpl{UserUsecase: userUsecase}
}

func (authController *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var userLoginRequest user.LoginRequestUser
	helper.ReadFromRequestBody(r, &userLoginRequest)

	userAuthData := authController.UserUsecase.Login(r.Context(), &userLoginRequest)

	expiredTokenTime := time.Now().Add(time.Minute * 5)
	jwtClaims := &config.JWTClaims{
		Email: userAuthData.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "laundry-app",
			ExpiresAt: jwt.NewNumericDate(expiredTokenTime),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	generatedToken, err := tokenAlgorithm.SignedString(config.JwtSecretKey)
	helper.PanicIfError(err)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    generatedToken,
		Path:     "/",
		HttpOnly: true,
	})

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success",
		Data:    "Login Success",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (authController *AuthControllerImpl) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Logout",
		Data:    "Success Logout",
	}
	helper.WriteToResponseBody(w, webResponse)
}
