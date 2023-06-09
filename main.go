package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/labadiejoaco/gin-gorm-api/database"
	"github.com/labadiejoaco/gin-gorm-api/router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
}

func main() {
	r := router.InitRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
