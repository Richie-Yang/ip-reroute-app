package main

import (
	"ip-reroute-app/handler"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(cors.Default())
	handler.Register(router)
	router.Run(":" + port) // listens on 0.0.0.0:8080 by default
}

func loadEnv() {
	env := os.Getenv("ENV_MODE")
	if env == "" {
		err := godotenv.Load("env/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
