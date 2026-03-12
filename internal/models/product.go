package models

import "time"

type Product struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID uint      `gorm:"not null" json:"category_id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock      int       `gorm:"default:0" json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationship - one product belongs to one category
	Category Category `gorm:"foreignKey:CategoryId" json:"category,omitempty"`
}

// Request structs
type CreateProductRequest struct {
	CategoryID uint    `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}

type UpdateProductRequest struct {
	CategoryID uint    `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}
