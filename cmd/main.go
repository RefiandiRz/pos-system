package main

import (
	"log"
	"os"

	"github.com/RefiandiRz/pos-system/config"
	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  No .env file found, using system environment variables")
	}

	// connect to database
	config.ConnectDatabase()

	// Auto migrate — creates tables if they don't exist
	config.MigrateDB(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Transaction{},
		&models.TransactionItem{},
	)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "POS System v1.0",
	})

	// Global middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup routes
	routes.SetupRoutes(app)

	// 404 handler
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}
