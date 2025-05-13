package entity

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	jwt.StandardClaims
}

func (CustomClaims) Valid() error {
	return nil
}