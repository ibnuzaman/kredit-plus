package service

import (
	"context"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
)

type customerService struct {
	repo      repository.CustomerRepository
	exception exception.Exception
}

type CustomerService interface {
	Information(ctx context.Context, customerId uint) *model.Customer
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		repo:      repo,
		exception: exception.NewException(),
	}
}

func (s *customerService) Information(ctx context.Context, customerId uint) *model.Customer {
	customer, err := s.repo.FindById(ctx, customerId)
	s.exception.Error(err)
	return customer
}
