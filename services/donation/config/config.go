package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimeZone string

	CampaignServiceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file, reading from environment")
	}

	return &Config{
		Port: envOr("PORT", "3003"),

		DBHost:     mustEnv("DB_HOST"),
		DBPort:     envOr("DB_PORT", "5432"),
		DBUser:     mustEnv("DB_USER"),
		DBPassword: mustEnv("DB_PASSWORD"),
		DBName:     mustEnv("DB_NAME"),
		DBSSLMode:  envOr("DB_SSLMODE", "disable"),
		DBTimeZone: envOr("DB_TIMEZONE", "Asia/Jakarta"),

		CampaignServiceURL: envOr("CAMPAIGN_SERVICE_URL", "http://campaign-service:3002"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode, c.DBTimeZone,
	)
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("config: required environment variable %s is not set", key)
	}
	return v
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
