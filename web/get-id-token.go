package web

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GetIdToken(bearerToken string) (string, error) {
	token, _ := jwt.Parse(bearerToken[7:], nil)

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("Error getting token from jwt")
	}

	userID, ok := claims["userID"].(string)

	if !ok {
		return "", errors.New("Error getting token from jwt CLAIMS")
	}

	return userID, nil
}
