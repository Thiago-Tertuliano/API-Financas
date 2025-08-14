package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	JWTSecret string
	DatabaseURL string
}

func LoadConfig() *Config {
	// Carrega as variáveis do .env se existir
	_ = godotenv.Load()

	cfg := &Config{
		Port:    getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "default"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/financa"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
