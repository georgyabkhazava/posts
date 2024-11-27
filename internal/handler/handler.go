package handler

import (
	"context"
	"github.com/georgyabkhazava/posts/internal/domain"
	"github.com/gin-gonic/gin"
)

type RegistrationService interface {
	RegistrationUser(ctx context.Context, name string, password string, email string) (int64, error)
	LoginUser(ctx context.Context, name string, password string, email string) (string, error)
}

type TwitService interface {
	CreateTwit(ctx context.Context, title string, text string, userId int64) (int64, error)
	GetTwitsByUserId(ctx context.Context, userId int64) ([]domain.Twit, error)
	DeleteTwitById(ctx context.Context, id int64, userId int64) error
	GetTwitById(ctx context.Context, id int64) (domain.Twit, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, userId int64, text string, twitId int64) (int64, error)
}

type Handler struct {
	service        RegistrationService
	twitservice    TwitService
	commentservice CommentService
}

// принимает аргумент service и создает объект хэндлера
func New(s RegistrationService, t TwitService, c CommentService) *Handler {
	return &Handler{
		service:        s,
		twitservice:    t,
		commentservice: c,
	}
}

// проверяет рабоспособность сервиса
func (h *Handler) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
