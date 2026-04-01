package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/raudhra/movie-reservation-system/config"
	"github.com/raudhra/movie-reservation-system/routes"
)

func main() {
	godotenv.Load()
	port := os.Getenv("APP_PORT")
	config.Connect()
	router := routes.SetupRoutes()
	router.Run(":" + port)
}
