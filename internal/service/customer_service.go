package service

import (
	"context"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
)

type customerService struct {
	repo      repository.CustomerRepository
	tenorRepo repository.TenorRepository
	exception exception.Exception
}

type CustomerService interface {
	Information(ctx context.Context, customerId uint) model.CustomerResponse
	Tenor(ctx context.Context, customerId uint) []model.TenorResponse
}

func NewCustomerService(repo repository.CustomerRepository, tenorRepo repository.TenorRepository) CustomerService {
	return &customerService{
		repo:      repo,
		tenorRepo: tenorRepo,
		exception: exception.NewException(),
	}
}

func (s *customerService) Information(ctx context.Context, customerId uint) model.CustomerResponse {
	customer, err := s.repo.FindById(ctx, customerId)
	s.exception.Error(err)
	return customer.ToResponse()
}

func (s *customerService) Tenor(ctx context.Context, customerId uint) []model.TenorResponse {
	tenors, err := s.tenorRepo.FindByCustomerId(ctx, customerId)
	s.exception.Error(err)

	res := []model.TenorResponse{}
	for _, tenor := range tenors {
		res = append(res, tenor.ToResponse())
	}
	return res
}
