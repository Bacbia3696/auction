package token

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Name string `json:"Name"`
	Id   int32    `json:"Id"`
	jwt.StandardClaims
}
var jwtKey = []byte("abcdefghijklmnopq") //config.Get().TokenSignKey

func GenToken(user db.User) (string, error) {
	expirationTime := time.Now().Add(120 * time.Second)
	claims := &Claims{
		Name: user.UserName,
		Id:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}



func Verify(tokenString string) (*Claims, error) {
	tk := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token.Valid {
		return token.Claims.(*Claims), nil
	}
	return nil, err
}
