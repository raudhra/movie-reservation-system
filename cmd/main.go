package main

import (
	"os"

	"github.com/raudhra/movie-reservation-system/config"
	"github.com/raudhra/movie-reservation-system/routes"
)

func main() {
	port := os.Getenv("APP_PORT")
	config.Connect()
	router := routes.SetupRoutes()
	router.Run(":" + port)
}
