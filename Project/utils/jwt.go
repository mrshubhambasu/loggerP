package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key_for_jwt")

type claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stoken, err := token.SignedString(jwtKey)

	return stoken, err
}

func ValidateJWT(tokenStr string) (*claims, error) {
	claims := &claims{}

	f := func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, claims, f)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
