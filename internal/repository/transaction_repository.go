package repository

import (
	"context"
	"kredit-plus/internal/model"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

type TransactionRepository interface {
	FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Transaction, error) {
	var transactions []model.Transaction

	limit, offset := limitOffset(page, perPage)
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Offset(offset).Limit(limit).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
