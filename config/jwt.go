package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("abhybababub")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
