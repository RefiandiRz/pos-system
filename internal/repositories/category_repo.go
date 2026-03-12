package repositories

import (
	"github.com/RefiandiRz/pos-system/config"
	"github.com/RefiandiRz/pos-system/internal/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	return categories, result.Error
}

func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	result := config.DB.First(&category, id)
	return &category, result.Error
}

func CreateCategory(category *models.Category) error {
	return config.DB.Create(category).Error
}

func UpdateCategory(category *models.Category) error {
	return config.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	return config.DB.Delete(&models.Category{}, id).Error
}

func CategoryExist(id uint) bool {
	var count int64
	config.DB.Model(&models.Category{}).Where("id = ?", id).Count(&count)
	return count > 0
}
