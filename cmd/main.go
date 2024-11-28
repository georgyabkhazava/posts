package main

import (
	"database/sql"
	"fmt"
	"github.com/georgyabkhazava/posts/internal/handler"
	"github.com/georgyabkhazava/posts/internal/middlewares"
	"github.com/georgyabkhazava/posts/internal/service/comment"
	"github.com/georgyabkhazava/posts/internal/service/registration"
	"github.com/georgyabkhazava/posts/internal/service/twit"
	"github.com/georgyabkhazava/posts/internal/service/verification_email"
	commentDB "github.com/georgyabkhazava/posts/internal/storage/comment"
	registrationDB "github.com/georgyabkhazava/posts/internal/storage/registration"
	twitDB "github.com/georgyabkhazava/posts/internal/storage/twit"
	verificationDB "github.com/georgyabkhazava/posts/internal/storage/verification"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic("no .env file found")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	registrationStorage := registrationDB.New(db)
	verificationStorage := verificationDB.New(db)
	twitStorage := twitDB.New(db)
	commentStorage := commentDB.New(db)

	verificationService := verification_email.New(verificationStorage, registrationStorage)

	registrationService := registration.New(registrationStorage, verificationService)
	twitService := twit.New(twitStorage)
	commentService := comment.New(twitStorage, commentStorage)

	h := handler.New(registrationService, twitService, commentService)
	middleware := middlewares.New()

	r.GET("/ping", h.HandlePing)
	r.POST("/registration", h.HandleRegistration)
	r.POST("/login", h.HandleLogin)
	r.POST("/twits/create", middleware.CheckToken, h.HandleTwit)
	r.POST("/twits/delete", middleware.CheckToken, h.HandleDeleteTwit)
	r.GET("/twits", middleware.CheckToken, h.HandleGetTwits)
	r.GET("/twits/:id", middleware.CheckToken, h.HandleGetTwit)
	r.POST("/comments/create", middleware.CheckToken, h.HandleComment)
	r.Run(":80")

}
