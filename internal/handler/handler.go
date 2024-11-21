package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type RegistrationService interface {
	RegistrationUser(ctx context.Context, name string, password string) (int64, error)
}

type Handler struct {
	service RegistrationService
}

func New(s RegistrationService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
