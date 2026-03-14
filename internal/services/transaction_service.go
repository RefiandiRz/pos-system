package services

import (
	"errors"
	"fmt"

	"github.com/RefiandiRz/pos-system/config"
	"github.com/RefiandiRz/pos-system/internal/models"
	"github.com/RefiandiRz/pos-system/internal/repositories"
)

func CreateTransaction(userID uint, req models.CreateTransactionRequest) (*models.Transaction, error) {
	// Validate items
	if len(req.Items) == 0 {
		return nil, errors.New("transaction must have at least one item")
	}

	// Validate payment method
	validMethods := map[models.PaymentMethod]bool{
		models.PaymentCash:   true,
		models.PaymentDebit:  true,
		models.PaymentCredit: true,
		models.PaymentQRIS:   true,
	}

	if !validMethods[req.PaymentMethod] {
		return nil, errors.New("invalid payment method. Use: cash, debit, credit, or qris")
	}

	// Start DB transaction
	// If any step fails, we can rollback to maintain data integrity
	var createdTransaction *models.Transaction

	err := config.DB.Transaction(func(tx *config.DBTransaction) error {
		var totalAmount float64
		var transactionItems []models.TransactionItem

		// Process each item
		for _, item := range req.Items {
			if item.Quantity <= 0 {
				return fmt.Errorf("quantity for product ID %d must be greater than zero", item.ProductID)
			}

			// fetch product
			product, err := repositories.GetProductByID(item.ProductID)
			if err != nil {
				return fmt.Errorf("product with ID %d not found", item.ProductID)
			}

			//  check stock availability
			if product.Stock < item.Quantity {
				return fmt.Errorf("insufficient stock fo '%s'. Available: %d, Requested: %d", product.Name, product.Stock, item.Quantity)
			}

			// decut stock
			if err := tx.Model(&models.Product{}).Where("id = ?", product.ID).UpdateColumn("stock", product.Stock-item.Quantity).Error; err != nil {
				return fmt.Errorf("failed to update stock for product '%s'", product.Name)
			}

			// Calculate subtotal
			subtotal := product.Price * float64(item.Quantity)
			totalAmount += subtotal

			transactionItems = append(transactionItems, models.TransactionItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				UnitPrice: product.Price,
				Subtotal:  subtotal,
			})
		}

		// Validate amount paid
		if req.AmountPaid < totalAmount {
			return fmt.Errorf("amount paid (%.2f) is less than total amount (%.2f)", req.AmountPaid, totalAmount)
		}

		// Calculate change
		change := req.AmountPaid - totalAmount

		// Create transaction record
		transaction := &models.Transaction{
			UserID:        userID,
			TotalAmount:   totalAmount,
			PaymentMethod: req.PaymentMethod,
			AmountPaid:    req.AmountPaid,
			Change:        change,
			Items:         transactionItems,
		}

		if err := tx.Create(transaction).Error; err != nil {
			return errors.New("failed to save transaction")
		}

		createdTransaction = transaction
		return nil
	})

	// end DB transaction
	if err != nil {
		return nil, err
	}

	// return full transaction with all relations preloaded
	return repositories.GetTransactionByID(createdTransaction.ID)
}

func GetAllTransaction() ([]models.Transaction, error) {
	return repositories.GetAllTransaction()
}

func GetTransactionByID(id uint) (*models.Transaction, error) {
	transaction, err := repositories.GetTransactionByID(id)
	if err != nil {
		return nil, fmt.Errorf("transaction with ID %d not found", id)
	}
	return transaction, nil
}
