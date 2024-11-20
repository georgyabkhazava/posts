package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RegistrationRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *Handler) HandleRegistration(c *gin.Context) {
	var request RegistrationRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// написать метод валидации, провалидировать и если что вернуть 400

	fmt.Println(request.Name, request.Password)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
