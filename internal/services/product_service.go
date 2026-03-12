package services

import (
	"errors"

	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/internal/repositories"
)

func GetAllProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func GetProductByID(id uint) (*models.Product, error) {
	product, err := repositories.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func CreateProduct(req models.CreateProductRequest) (*models.Product, error) {
	if req.Name == "" {
		return nil, errors.New("product name is required")
	}

	if req.Price <= 0 {
		return nil, errors.New("product price must be greater than zero")
	}

	if req.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	// check if category exists
	if !repositories.CategoryExist(req.CategoryID) {
		return nil, errors.New("category not found")
	}

	product := &models.Product{
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	if err := repositories.CreateProduct(product); err != nil {
		return nil, errors.New("failed to create product, name may already exist")
	}

	//  Return with category preloaded
	return repositories.GetProductByID(product.ID)
}

func UpdateProduct(id uint, req models.UpdateProductRequest) (*models.Product, error) {
	product, err := repositories.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// validate fields
	if req.Name == "" {
		return nil, errors.New("product name is required")
	}

	if req.Price <= 0 {
		return nil, errors.New("product price must be greater than zero")
	}

	if req.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	// check if category exists
	if !repositories.CategoryExist(req.CategoryID) {
		return nil, errors.New("category not found")
	}

	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock

	if err := repositories.UpdateProduct(product); err != nil {
		return nil, errors.New("failed to update product")
	}

	return repositories.GetProductByID(product.ID)
}

func DeleteProduct(id uint) error {
	if err := repositories.DeleteProduct(id); err != nil {
		return errors.New("product not found")
	}

	return repositories.DeleteProduct(id)
}
