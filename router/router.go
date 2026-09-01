package router

import (
	"donasi-app/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, userHandler *handler.UserHandler) {
	api := e.Group("/api/v1")

	// Endpoint User (Ganti api.Post jadi api.POST)
	api.POST("/register", userHandler.Register)
}
