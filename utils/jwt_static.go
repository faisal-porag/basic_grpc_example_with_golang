package utils

import "github.com/golang-jwt/jwt/v4"

type ServiceAuthClaims struct {
	jwt.RegisteredClaims
}

var JwtSecretKey = []byte("$%^HSBS@#fg$5HJ^dshfsd657hb%45^5b")
