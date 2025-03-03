package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV *Config

func LoadConfig() {
	// Load .env only in local environment
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found, relying on system environment variables")
		}
	}

	ENV = &Config{
		PORT:        os.Getenv("PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}
}
