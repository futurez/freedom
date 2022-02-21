package futils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

var jwtkey = []byte("www.freedom.com")

func GetUserToken(userId uint64) (string, error) {
	expireTime := time.Now().Add(7*24*time.Hour)

	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:  "freedom",
			Subject:  "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func parseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, claims, err
}

func CheckUserToken(userId uint64, strToken string) bool {
	token, claims, err := parseToken(strToken)
	if err != nil || !token.Valid {
		return false
	}
	if claims.UserId == userId {
		return true
	}
	return false
}