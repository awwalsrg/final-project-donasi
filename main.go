package main

import (
	"log"
	"os"

	"donasi-app/config"
	"donasi-app/handler"
	"donasi-app/repositories"
	"donasi-app/router"
	"donasi-app/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("[Warning] Error loading .env file, falling back to system environment variables")
	}

	// 2. Inisialisasi Koneksi Database (mengambil nama database dari env)
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "donasi_db" // Nama database default jika env tidak ada
	}

	db := config.ConnectDB(dbName)
	if db == nil {
		log.Fatal("[Fatal] Failed to connect to database. Application shutting down...")
	}

	// 3. Inisialisasi Framework Echo
	e := echo.New()

	// 4. Konfigurasi Middleware Global
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} | ${method} | ${uri} (${latency_human})\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// 5. Dependency Injection (Wiring Layers: Repository -> Usecase -> Handler)
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// 6. Daftarkan Routes API
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  "success",
			"message": "Donasi App API is Running smoothly!",
		})
	})

	// Daftarkan group route API v1
	router.InitRouter(e, userHandler)

	// 7. Jalankan Server pada Port yang ditentukan
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("[Server] Starting server on port %s...", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("[Fatal] Shutting down the server: %v", err)
	}
}
