package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	ID          string
	Email       string
	PhoneNumber string
	jwt.StandardClaims
}

func GenerateToken(id string, email string, phone_number string) (signed string, err error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	JWT_LIFE_INT, err := strconv.Atoi(os.Getenv("JWT_LIFE"))
	if err != nil {
		log.Panic(err)
	}

	claims := &SignedDetails{
		ID:          id,
		Email:       email,
		PhoneNumber: phone_number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(JWT_LIFE_INT).Unix()),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	return token, err
}
