package main

import (
	"donasi-app/config"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Inisialisasi Database
	config.ConnectDB()

	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route Dasar
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Donasi App API is Running!")
	})

	// Jalankan Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
