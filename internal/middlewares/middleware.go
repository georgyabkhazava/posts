package middlewares

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/georgyabkhazava/posts/internal/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const authorizationHeader = "Authorization"

func New() *Middleware {
	return &Middleware{}
}

type Middleware struct{}

func (h *Middleware) CheckToken(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}
	c.Set("claims", claims)
}

func (h *Middleware) parseAuthHeader(c *gin.Context) (domain.AuthClaims, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return domain.AuthClaims{}, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return domain.AuthClaims{}, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return domain.AuthClaims{}, errors.New("token is empty")
	}

	claims, err := h.parse(headerParts[1])
	if err != nil {
		return domain.AuthClaims{}, err
	}

	return claims, nil
}

func (h *Middleware) parse(accessToken string) (domain.AuthClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected string method")
		}
		return []byte(domain.Key), nil
	})
	if err != nil {
		return domain.AuthClaims{}, err
	}

	b, _ := json.Marshal(token.Claims)

	var claims domain.AuthClaims
	json.Unmarshal(b, &claims)

	return claims, nil
}
