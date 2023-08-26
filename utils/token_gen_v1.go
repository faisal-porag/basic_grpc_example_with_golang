package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"log"
	"time"
)

var (
	privateKey []byte
	publicKey  []byte
)

func init() {
	// Load private key
	privKeyBytes, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	privateKey = privKeyBytes

	// Load public key
	pubKeyBytes, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	publicKey = pubKeyBytes
}

func GenerateTokenForServiceToServiceCall() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	privateKeyObj, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(privateKeyObj)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateServiceToServiceCallToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
