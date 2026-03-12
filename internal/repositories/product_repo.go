package repositories

import (
	"github.com/RefiandiRz/pos-system/config"
	"github.com/RefiandiRz/pos-system/internal/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	// Preload category data alongside each product
	result := config.DB.Preload("Category").Find(&products)
	return products, result.Error
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := config.DB.Preload("Category").First(&product, id)
	return &product, result.Error
}

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}

// DeductStock reduces product stock by a given quantity
// Used during checkout in Phase 4
func DeductStock(productID uint, quantity int) error {
	return config.DB.Model(&models.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		UpdateColumn("stock", config.DB.Raw("stock - ?", quantity)).
		Error
}
