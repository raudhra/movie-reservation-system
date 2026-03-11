package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	DatabaseURL string
	App_Port    string
	App_Env     string
}

func LoadConfig() *config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the env file")
	}

	return &config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:2645@localhost:5432/movie-reservation-system?sslmode=disable")
		App_Port: getEnv("APP_PORT", "8080")
		App_Env: getEnv("APP_ENV", "development")
	}
}

func Connect(){
	LoadConfig()
	db, err:= gorm.Open("postgres", config.DatabaseURL)
	if err != nil{
		log.Fatal("Unable to Connect To Database")
	}

	log.Println("Databse connected successfully")
}
func getDB() *gorm.DB{
	return db
}
func getEnv(key string, defaultValue string) string{
	value := os.Getenv(key)
	if value == ""{
		return defaultValue
	}
	return value
}
