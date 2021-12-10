package utils

import (
	"admire-avatar/entities"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func GetJWTClaims(tokenData string, key []byte) (*entities.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenData, &entities.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in token parsing")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entities.UserClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
