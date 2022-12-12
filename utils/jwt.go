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
	LoadConfig()
	SECRET_KEY := Config.SecretKey
	JWT_LIFE_INT, err := strconv.Atoi(Config.JwtLife)
	if err != nil {
		log.Panic(err)
	}

	claims := &SignedDetails{
		ID:          id,
		Email:       email,
		PhoneNumber: phone_number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(JWT_LIFE_INT)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}
	// token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "The token is expired"
		return
	}
	return claims, msg
}

func Auth(signedToken string) *SignedDetails {
	claims, msg := ValidateToken(signedToken)
	if msg != "" {
		log.Fatalf(msg)
	}
	return claims
}
