package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"unicode"
)

type RegistrationRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// isValidatePassword метод валидации пароля, проверяет соответствует ли пароль заданным параметрам
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

func (r RegistrationRequest) isValidateEmail() bool {
	if len(r.Email) > 255 {
		return false
	}

	if strings.Index(r.Email, "@") != -1 {
		return true
	}

	return false
}

// HandleRegistration метод регистрации
// ррр
func (h *Handler) HandleRegistration(c *gin.Context) {
	var request RegistrationRequest
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

	if request.isValidateEmail() == false {
		c.JSON(400, gin.H{
			"message": "Not Validate",
		})
		return
	}

	id, err := h.service.RegistrationUser(c, request.Name, request.Password, request.Email)
	if err != nil {
		println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}
