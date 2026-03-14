package repositories

import (
	"github.com/RefiandiRz/pos-system/config"
	"github.com/RefiandiRz/pos-system/internal/models"
)

func CreateTransaction(tx *models.Transaction) error {
	return config.DB.Create(tx).Error
}

func GetAllTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	result := config.DB.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		Order("created_at desc").
		Find(&transaction)
	return transaction, result.Error
}

func GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := config.DB.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		First(&transaction, id)
	return &transaction, result.Error
}

// GetProductWithLock fetches a product and locks the row for update
// This prevents race conditions when two checkouts happen simultaneously
func GetProductWithLock(productID uint) (*models.Product, error) {
	var product models.Product
	result := config.DB.Set("gorm:query_option", "FOR UPDATE").First(&product, productID)
	return &product, result.Error
}
