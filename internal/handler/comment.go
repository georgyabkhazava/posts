package handler

import (
	"github.com/georgyabkhazava/posts/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

type CreateComment struct {
	Text   string
	TwitId int64
}

func (c CreateComment) isValidateComment() bool {
	if c.Text == "" {
		return false
	}

	if utf8.RuneCountInString(c.Text) > 1000 {
		return false
	}
	return true
}

func (h *Handler) HandleComment(c *gin.Context) {
	var request CreateComment
	err := c.Bind(&request) // метод Bind достает данные из запроса и засовывает их в структуру
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.isValidateComment() == false {
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

	id, err := h.commentservice.CreateComment(c, claims.UserID, request.Text, request.TwitId)
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
