package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserID int64
	jwt.StandardClaims
}

func GenToken(ID int64) (string, error) {
	mySigningKey := []byte("bluemsun")
	c := MyClaims{
		UserID: ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    "qza",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseToken(tokenString string) (int64, error) {
	mySigningKey := []byte("bluemsun")
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return 0, err
	}
	return token.Claims.(*MyClaims).UserID, err
}

func TokenCheck(tokenString string, id int64) int {
	return 0
}
