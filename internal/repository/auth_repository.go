package repository

import (
	"context"
	"kredit-plus/internal/model"
	"kredit-plus/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type authRepository struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.Customer, error)
	FindById(ctx context.Context, id int) (*model.Customer, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db:  db,
		log: logger.Get("auth_repository"),
	}
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&customer).Error
	if err != nil {
		r.log.Error().Err(err).Msg("failed to find customer by email")
		return nil, err
	}
	return &customer, nil
}

func (r *authRepository) FindById(ctx context.Context, id int) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		r.log.Error().Err(err).Msg("failed to find customer by id")
		return nil, err
	}
	return &customer, nil
}
