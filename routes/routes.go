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

	// Categories (admin only for write)
	protected.Get("/categories", handlers.GetAllCategories)
	protected.Get("/categories/:id", handlers.GetCategoryByID)
	protected.Post("/categories", middleware.AdminOnly, handlers.CreateCategory)
	protected.Put("/categories/:id", middleware.AdminOnly, handlers.UpdateCategory)
	protected.Delete("/categories/:id", middleware.AdminOnly, handlers.DeleteCategory)

	// Products (admin only for write)
	protected.Get("/products", handlers.GetAllProducts)
	protected.Get("/products/:id", handlers.GetProductByID)
	protected.Post("/products", middleware.AdminOnly, handlers.CreateProduct)
	protected.Put("/products/:id", middleware.AdminOnly, handlers.UpdateProduct)
	protected.Delete("/products/:id", middleware.AdminOnly, handlers.DeleteProduct)

}
