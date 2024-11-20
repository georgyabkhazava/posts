package main

import (
	"github.com/georgyabkhazava/posts/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	h := handler.New()

	r.GET("/ping", h.HandlePing)
	r.POST("/registration", h.HandleRegistration)

	r.Run(":80")
}
