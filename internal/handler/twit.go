package handler

import (
	"github.com/georgyabkhazava/posts/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

type CreateTwit struct {
	Title string
	Text  string
}

func (c CreateTwit) isValidateTwit() bool {
	if c.Text == "" || c.Title == "" {
		return false
	}

	if utf8.RuneCountInString(c.Text) > 1000 || utf8.RuneCountInString(c.Title) > 30 {
		return false
	}
	return true
}

func (h *Handler) HandleTwit(c *gin.Context) {
	var request CreateTwit
	err := c.Bind(&request) // метод Bind достает данные из запроса и засовывает их в структуру
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.isValidateTwit() == false {
		c.JSON(400, gin.H{
			"message": "Not Validate",
		})
		return
	}

	claims, err := domain.GetClaims(c)
	if err != nil {
		println(err.Error())
		return
	}

	id, err := h.twitservice.CreateTwit(c, request.Text, request.Title, claims.UserID)
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
