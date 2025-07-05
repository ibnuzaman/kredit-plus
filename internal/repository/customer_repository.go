package repository

import (
	"context"
	"kredit-plus/internal/model"
	"kredit-plus/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type customerRepository struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type CustomerRepository interface {
	FindById(ctx context.Context, id uint) (*model.Customer, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db:  db,
		log: logger.Get("customer_repository"),
	}
}

func (r *customerRepository) FindById(ctx context.Context, id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		r.log.Error().Err(err).Uint("id", id).Msg("failed to find customer by id")
		return nil, err
	}
	return &customer, nil
}
