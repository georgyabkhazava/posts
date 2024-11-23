package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (l Login) isValidatePassword() bool {
	if len(l.Password) < 8 {
		return false
	}
	for _, p := range l.Password {
		if unicode.IsUpper(p) {
			return true
		}
	}
	return false
}

func (h *Handler) HandleLogin(c *gin.Context) {
	var request Login
	err := c.Bind(&request) // метод Bind достает данные из запроса и засовывает их в структуру
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

	token, err := h.service.LoginUser(c, request.Name, request.Password)
	if err != nil {
		println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
