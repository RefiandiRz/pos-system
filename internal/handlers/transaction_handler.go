package handlers

import (
	"strconv"

	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/internal/services"
	"github.com/gofiber/fiber/v3"
)

// POST /api/transactions
func CreateTransaction(c fiber.Ctx) error {
	var req models.CreateTransactionRequest

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request body",
		})
	}

	// Get userID from JWT claims set by middleware
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "unauthorized",
		})
	}

	transcation, err := services.CreateTransaction(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "transaction created successfully",
		"data":    transcation,
	})
}

func GetAllTransaction(c fiber.Ctx) error {
	transactions, err := services.GetAllTransaction()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to retrieve transactions",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "transactions retrieved successfully",
		"data":    transactions,
	})
}

func GetTransactionByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid transaction ID",
		})
	}

	transaction, err := services.GetTransactionByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "transaction retrieved successfully",
		"data":    transaction,
	})
}
