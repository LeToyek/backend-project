package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	UserID string
	jwt.StandardClaims
}

const SECRET_KEY = "super duper secret omg"

func GenerateTokens(userID string) (signedToken string, signedRefreshToken string, err error) {
	claims := &Credentials{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 5).Unix(),
		},
	}

	refreshClaims := &Credentials{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))

	return token, refreshToken, err
}

func ValidateToken(signedToken string) *Credentials {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&Credentials{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(*Credentials)
	if !ok {
		panic(err)
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		panic(err)
	}
	return claims
}
