package config

import (
	"fmt"
	"log"
	"os"
	"time"

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

	JWTSecret string
	JWTExpiry time.Duration

	CampaignServiceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file, reading from environment")
	}

	return &Config{
		Port: envOr("PORT", "3001"),

		DBHost:     mustEnv("DB_HOST"),
		DBPort:     envOr("DB_PORT", "5432"),
		DBUser:     mustEnv("DB_USER"),
		DBPassword: mustEnv("DB_PASSWORD"),
		DBName:     mustEnv("DB_NAME"),
		DBSSLMode:  envOr("DB_SSLMODE", "disable"),
		DBTimeZone: envOr("DB_TIMEZONE", "Asia/Jakarta"),

		JWTSecret: mustEnv("JWT_SECRET"),
		JWTExpiry: envDuration("JWT_EXPIRY", 24*time.Hour),

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

func envDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("config: bad %s=%q, using %s", key, v, fallback)
		return fallback
	}
	return d
}
