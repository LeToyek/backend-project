package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	UserID string
	jwt.StandardClaims
}

var SECRET_KEY = []byte("super_duper_secret_omg")

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(SECRET_KEY)

	if err != nil {
		panic(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	return tokenStr, refreshToken, err
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
