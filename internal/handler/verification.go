package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerificationCode struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
}

func (v VerificationCode) isValidateCode() bool {
	if len(v.Code) < 4 {
		return false
	}
	return true

}
func (h *Handler) VerificationCodeHandler(c *gin.Context) {
	var request VerificationCode
	err := c.Bind(&request) // метод Bind достает данные из запроса и засовывает их в структуру
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.isValidateCode() == false {
		c.JSON(400, gin.H{
			"message": "Not Validate",
		})
		return
	}

}
