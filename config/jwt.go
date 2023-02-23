package config

import "github.com/golang-jwt/jwt/v4"

var JwtSecretKey = []byte("golangLaundryAppJWTSecretKey")

type JWTClaims struct {
	Email string
	jwt.RegisteredClaims
}
