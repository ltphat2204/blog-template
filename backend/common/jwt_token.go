package common

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(data map[string]interface{}) (string, error) {
	expireTime := time.Now().Add(30 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data["username"],
		"exp":      expireTime,
	})

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
