package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type RegistrationService interface {
	RegistrationUser(ctx context.Context, name string, password string) (int64, error)
	LoginUser(ctx context.Context, name string, password string) (string, error)
}

type TwitService interface {
	CreateTwit(ctx context.Context, title string, text string, userId int64) (int64, error)
}

type Handler struct {
	service     RegistrationService
	twitservice TwitService
}

// принимает аргумент service и создает объект хэндлера
func New(s RegistrationService, t TwitService) *Handler {
	return &Handler{
		service:     s,
		twitservice: t,
	}
}

// проверяет рабоспособность сервиса
func (h *Handler) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
