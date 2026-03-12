package services

import (
	"errors"

	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/internal/repositories"
)

func GetAllCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func GetCategoryById(id uint) (*models.Category, error) {
	category, err := repositories.GetCategoryByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {
	if req.Name == "" {
		return nil, errors.New("category name is required")
	}

	category := &models.Category{
		Name: req.Name,
	}

	if err := repositories.CreateCategory(category); err != nil {
		return nil, errors.New("failed to create category, name may already exist")
	}

	return category, nil
}

func UpdateCategory(id uint, req models.UpdateCategoryRequest) (*models.Category, error) {
	category, err := repositories.GetCategoryByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	if req.Name == "" {
		return nil, errors.New("category name is required")

	}

	category.Name = req.Name

	if err := repositories.UpdateCategory(category); err != nil {
		return nil, errors.New("failed to update category")
	}

	return category, nil
}

func DeleteCategory(id uint) error {
	if !repositories.CategoryExist(id) {
		return errors.New("category not found")
	}
	return repositories.DeleteCategory(id)
}
