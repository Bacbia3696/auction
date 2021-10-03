package token

import (
	"time"

	"github.com/bacbia3696/auction/internal/config"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func Create(username string, duration time.Duration) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(config.Get().TokenSignKey)
}

func Verify(tokenString string) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(config.Get().TokenSignKey), nil
	})
	if token.Valid {
		return token.Claims.(*Claims), nil
	}
	return nil, err
}
