package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"unicode"
)

type RegistrationRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r RegistrationRequest) isValidatePassword() bool {
	if len(r.Password) < 8 {
		return false
	}
	for _, p := range r.Password {
		if unicode.IsUpper(p) {
			return true
		}
	}
	return false
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

	if request.isValidatePassword() == false {
		c.JSON(400, gin.H{
			"message": "Not Validate",
		})
		return
	}

	// написать метод валидации, провалидировать и если что вернуть 400

	fmt.Println(request.Name, request.Password)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
