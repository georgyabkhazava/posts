package main

import (
	"database/sql"
	"fmt"
	"github.com/georgyabkhazava/posts/internal/handler"
	"github.com/georgyabkhazava/posts/internal/service"
	"github.com/georgyabkhazava/posts/internal/storage"
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

	registrationStorage := storage.New(db)

	registrationService := service.New(registrationStorage)

	h := handler.New(registrationService)

	r.GET("/ping", h.HandlePing)
	r.POST("/registration", h.HandleRegistration)

	r.Run(":80")
}
