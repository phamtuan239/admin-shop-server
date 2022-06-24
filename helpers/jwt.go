package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-Go"
)

type Claims struct {
	jwt.StandardClaims
}

func GenerateJWT(issuer string, expire int64, secret string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: expire,
	})

	return claims.SignedString([]byte(secret))
}

func ParseJWT(jwtCookie string, secret string) (*jwt.StandardClaims, error) {
	
	token, err := jwt.ParseWithClaims(jwtCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return &jwt.StandardClaims{}, err
	}

	if !token.Valid {
		return &jwt.StandardClaims{}, errors.New("token not valid")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}
