package main

import (
	"go-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.Static("/upload", "./upload")

	routes.Serve(r)
	uploadDir := [...]string{"photos", "users"}
	for _, p := range uploadDir {
		os.MkdirAll("/upload/"+p, 0755)
	}
	r.Run(":" + port)
}
