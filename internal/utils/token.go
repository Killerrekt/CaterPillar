package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(EmployeeID string, secretKey string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"EmployeeID": EmployeeID,
		"exp":        time.Now().Add(expiration).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
