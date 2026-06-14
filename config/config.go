package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}
	return Config{
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

}
