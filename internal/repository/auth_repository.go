package repository

import (
	"context"
	"kredit-plus/internal/model"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.Customer, error)
	FindById(ctx context.Context, id int) (*model.Customer, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *authRepository) FindById(ctx context.Context, id int) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
