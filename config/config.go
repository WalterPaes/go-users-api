package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	AppName    string
	AppPort    string
	DbName     string
	JwtSecret  string
	JwtExpires int
}

func Load() *config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExpires, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN_MINUTES"))

	return &config{
		AppName:    os.Getenv("APP_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		JwtExpires: jwtExpires,
	}
}
