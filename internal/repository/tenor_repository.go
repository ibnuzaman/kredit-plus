package repository

import (
	"context"
	"kredit-plus/internal/model"

	"gorm.io/gorm"
)

type tenorRepository struct {
	db *gorm.DB
}

type TenorRepository interface {
	FindByCustomerId(ctx context.Context, customerId uint) ([]model.Tenor, error)
}

func NewTenorRepository(db *gorm.DB) TenorRepository {
	return &tenorRepository{
		db: db,
	}
}

func (r *tenorRepository) FindByCustomerId(ctx context.Context, customerId uint) ([]model.Tenor, error) {
	var tenors []model.Tenor
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Find(&tenors).Error
	if err != nil {
		return nil, err
	}
	return tenors, nil
}
