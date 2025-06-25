package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLDSN string
	AppPort  string
}

func LoadConfig() *Config {
	// Загружаем .env (если есть)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN env variable is required")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		MySQLDSN: dsn,
		AppPort:  port,
	}
}
