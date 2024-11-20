package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
