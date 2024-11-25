package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	Key = "sagoooj24325j22sa0sd"
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

	return token.SignedString([]byte(Key))
}

func GetClaims(ctx *gin.Context) (AuthClaims, error) {
	claims, exists := ctx.Get("claims")
	if claims == nil {
		return AuthClaims{}, errors.New("claims is nil")
	}

	if !exists {
		return AuthClaims{}, errors.New("claims not exists")
	}

	return claims.(AuthClaims), nil
}
