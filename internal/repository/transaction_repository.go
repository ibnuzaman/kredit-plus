package repository

import (
	"context"
	"kredit-plus/internal/model"
	"kredit-plus/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type TransactionRepository interface {
	FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Transaction, error)
	FindByLoanId(ctx context.Context, loanId uint) ([]model.Transaction, error)
	Create(ctx context.Context, transaction *model.Transaction) error
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db:  db,
		log: logger.Get("transaction_repository"),
	}
}

func (r *transactionRepository) FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Transaction, error) {
	var transactions []model.Transaction

	limit, offset := limitOffset(page, perPage)
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Offset(offset).Limit(limit).Find(&transactions).Error
	if err != nil {
		r.log.Error().Err(err).Uint("customer_id", customerId).Uint("page", page).Uint("per_page", perPage).Msg("Failed to find transactions by customer ID")
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) FindByLoanId(ctx context.Context, loanId uint) ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := r.db.WithContext(ctx).Where("loan_id = ?", loanId).Find(&transactions).Error
	if err != nil {
		r.log.Error().Err(err).Uint("loan_id", loanId).Msg("Failed to find transactions by loan ID")
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			r.log.Error().Err(err).Msg("Failed to create transaction")
			return err
		}

		if err := tx.Model(&model.Loan{}).Where("id = ?", transaction.LoanID).Update("total_paid", gorm.Expr("total_paid + ?", 1)).Error; err != nil {
			r.log.Error().Err(err).Uint("loan_id", transaction.LoanID).Msg("Failed to update total paid for loan")
			return err
		}

		return nil
	})

	if err != nil {
		r.log.Error().Err(err).Msg("Failed to complete transaction creation")
		return err
	}

	return nil
}
