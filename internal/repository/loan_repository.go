package repository

import (
	"context"
	"kredit-plus/internal/model"
	"kredit-plus/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type loanRepository struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type LoanRepository interface {
	FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Loan, error)
	GetLastLoanByCustomerId(ctx context.Context, customerId uint) (*model.Loan, error)
	GetById(ctx context.Context, id uint) (*model.Loan, error)
	Create(ctx context.Context, loan *model.Loan) error
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{
		db:  db,
		log: logger.Get("loan_repository"),
	}
}

func (r *loanRepository) FindByCustomerId(ctx context.Context, customerId, page, perPage uint) ([]model.Loan, error) {
	var loans []model.Loan

	limit, offset := limitOffset(page, perPage)
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Offset(offset).Limit(limit).Find(&loans).Error
	if err != nil {
		r.log.Error().Err(err).Uint("customer_id", customerId).Msg("failed to find loans by customer id")
		return nil, err
	}

	return loans, nil
}

func (r *loanRepository) GetLastLoanByCustomerId(ctx context.Context, customerId uint) (*model.Loan, error) {
	var loan model.Loan
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Order("created_at DESC").Limit(1).First(&loan).Error
	if err != nil {
		r.log.Error().Err(err).Uint("customer_id", customerId).Msg("failed to get last loan by customer id")
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) GetById(ctx context.Context, id uint) (*model.Loan, error) {
	var loan model.Loan
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&loan).Error
	if err != nil {
		r.log.Error().Err(err).Uint("loan_id", id).Msg("failed to get loan by id")
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) Create(ctx context.Context, loan *model.Loan) error {
	err := r.db.WithContext(ctx).Create(loan).Error
	if err != nil {
		r.log.Error().Err(err).Msg("failed to create loan")
		return err
	}
	return nil
}
