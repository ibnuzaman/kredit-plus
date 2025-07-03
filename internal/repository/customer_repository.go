package repository

import (
	"context"
	"kredit-plus/internal/model"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

type CustomerRepository interface {
	FindById(ctx context.Context, id uint) (*model.Customer, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) FindById(ctx context.Context, id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
