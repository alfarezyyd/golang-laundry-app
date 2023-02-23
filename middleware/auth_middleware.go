package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"golang-laundry-app/config"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/web"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/login" {
			next.ServeHTTP(writer, request)
			return
		}

		cookieData, err := request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				webResponse := web.WebResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Data:    "Unauthorized",
				}
				helper.WriteToResponseBody(writer, webResponse)
			}
			return
		}
		tokenString := cookieData.Value
		jwtClaims := &config.JWTClaims{}
		parsedTokenJwt, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecretKey, nil
		})

		if err != nil {
			exceptionType, _ := err.(*jwt.ValidationError)
			switch exceptionType.Errors {
			case jwt.ValidationErrorExpired:
				writer.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Data:    "Unauthorized Token Expired",
				}
				helper.WriteToResponseBody(writer, webResponse)
				return
			default:
				writer.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized - Invalid Token",
					Data:    "Unauthorized - Invalid Token",
				}
				helper.WriteToResponseBody(writer, webResponse)
				return
			}
		}

		if !parsedTokenJwt.Valid {
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    "Unauthorized",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return

		}
		next.ServeHTTP(writer, request)
	})
}
