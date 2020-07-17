package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func GenerateJwtToken(userId uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Failed to load .env file")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["aud"] = os.Getenv("JWT_AUDIENCE")
	claims["exp"] = time.Now().Add(time.Minute * 36600).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
