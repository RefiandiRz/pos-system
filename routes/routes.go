package routes

import "github.com/gofiber/fiber/v3"

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

	// Auth routes
	_ = api
}
