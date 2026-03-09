package routes

import (
	"github.com/RefiandiRz/pos-system/internal/handlers"
	"github.com/RefiandiRz/pos-system/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// Health Check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "POS System is running!",
		})
	})

	// API group
	api := app.Group("/api")

	// Auth (public)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes (require authentication)
	protected := api.Group("", middleware.Protected)
	_ = protected
}
