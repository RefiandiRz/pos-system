package models

import "time"

type PaymentMethod string

const (
	PaymentCash   PaymentMethod = "cash"
	PaymentDebit  PaymentMethod = "debit"
	PaymentCredit PaymentMethod = "credit"
	PaymentQRIS   PaymentMethod = "qris"
)

// Transaction is the main order/sale record
type Transaction struct {
	ID            uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint          `gorm:"not null" json:"user_id"`
	TotalAmount   float64       `gorm:"type:decimal(12,2);not null" json:"total_amount"`
	PaymentMethod PaymentMethod `gorm:"type:varchar(20);not null" json:"payment_method"`
	AmountPaid    float64       `gorm:"type:decimal(12,2);not null" json:"amount_paid"`
	Change        float64       `gorm:"type:decimal(12,2);not null" json:"change"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`

	// Relationship
	User  User              `gorm:"foreignKey:UserID" json:"cashier,omitempty"`
	Items []TransactionItem `gorm:"foreignKey:TransactionID" json:"items,omitempty"`
}

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID uint    `gorm:"not null" json:"transaction_id"`
	ProductID     uint    `gorm:"not null" json:"product_id"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	UnitPrice     float64 `gorm:"type:decimal(12,2);not null" json:"unit_price"`
	Subtotal      float64 `gorm:"type:decimal(12,2);not null" json:"subtotal"`

	// Relationship
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// Request struct
type TransactionItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateTransactionRequest struct {
	PaymentMethod PaymentMethod            `json:"payment_method"`
	AmountPaid    float64                  `json:"amount_paid"`
	Items         []TransactionItemRequest `json:"items"`
}
