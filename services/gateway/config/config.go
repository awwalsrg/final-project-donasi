package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string

	UserServiceURL     string
	CampaignServiceURL string
	DonationServiceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file, reading from environment")
	}

	return &Config{
		Port:      envOr("PORT", "8080"),
		JWTSecret: mustEnv("JWT_SECRET"),

		UserServiceURL:     envOr("USER_SERVICE_URL", "http://user-service:3001"),
		CampaignServiceURL: envOr("CAMPAIGN_SERVICE_URL", "http://campaign-service:3002"),
		DonationServiceURL: envOr("DONATION_SERVICE_URL", "http://donation-service:3003"),
	}
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
