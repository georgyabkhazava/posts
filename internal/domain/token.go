package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	key = "sagoooj24325j22sa0sd"
)

var (
	ttl = time.Hour * 24
)

type AuthClaims struct {
	jwt.StandardClaims

	UserID int64
}

func GenerateAccessToken(userID int64) (tokenStr string, err error) {
	claims := AuthClaims{}
	claims.ExpiresAt = time.Now().Add(ttl).Unix()
	claims.UserID = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))
}

func DecodeToken(tokenStr string) (*AuthClaims, error) {
	claims := &AuthClaims{}
	jwt.
	if err != nil {}
}
