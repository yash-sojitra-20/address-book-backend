package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yash-sojitra-20/address-book-backend/internal/logger"
	"go.uber.org/zap"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JwtSecret string
	SMTPUser string
	SMTPPass string
	SMTPHost string
	SMTPPort string
}

func Load() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// log.Println("No .env file found, using system environment variables")
		logger.Logger.Warn("No .env file found, using system environment variables")

	}

	cfg := &Config{
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
		ServerPort: getEnv("SERVER_PORT"),
		JwtSecret:  getEnv("JWT_SECRET"),
		SMTPUser:  getEnv("SMTP_USER"),
		SMTPPass:  getEnv("SMTP_PASS"),
		SMTPHost:  getEnv("SMTP_HOST"),
		SMTPPort:  getEnv("SMTP_PORT"),
	}

	// log.Println("config loaded")
	logger.Logger.Info("config loaded")
	return cfg
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		// log.Fatalf("Environment variable %s is required", key)
		logger.Logger.Fatal("environment variable required", zap.String("key", key))

	}
	return value
}
