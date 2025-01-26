package main

import (
	"log"
	"os"

	"api-test/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	port := os.Getenv("PORT")

	routes.RegisterUserRoutes(router)

	err = router.Run(port)
	if err != nil {
		panic("Failed to start server on port " + port + " due to error: " + err.Error())
	}
}
