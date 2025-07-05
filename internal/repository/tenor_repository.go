package repository

import (
	"context"
	"kredit-plus/internal/model"
	"kredit-plus/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type tenorRepository struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type TenorRepository interface {
	FindByCustomerId(ctx context.Context, customerId uint) ([]model.Tenor, error)
}

func NewTenorRepository(db *gorm.DB) TenorRepository {
	return &tenorRepository{
		db:  db,
		log: logger.Get("tenor_repository"),
	}
}

func (r *tenorRepository) FindByCustomerId(ctx context.Context, customerId uint) ([]model.Tenor, error) {
	var tenors []model.Tenor
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerId).Find(&tenors).Error
	if err != nil {
		r.log.Error().Err(err).Uint("customer_id", customerId).Msg("failed to find tenors by customer id")
		return nil, err
	}
	return tenors, nil
}
