package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthClaims struct {
	UserId     uint
	Authorized bool
	*jwt.StandardClaims
}

func GenerateJwtToken(userId uint) (string, error) {
	aud := os.Getenv("JWT_AUDIENCE")
	if aud == "" {
		return "", errors.New("JWT_AUDIENCE can not be empty")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET can not be empty")
	}

	claims := &AuthClaims{
		UserId:     userId,
		Authorized: true,
		StandardClaims: &jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: time.Now().Add(time.Minute * 36600).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
