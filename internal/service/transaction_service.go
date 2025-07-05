package service

import (
	"context"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
)

type transactionService struct {
	repo      repository.TransactionRepository
	exception exception.Exception
}

type TransactionService interface {
	Transaction(ctx context.Context, customerId, page, perPage uint) []model.Transaction
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		repo:      repo,
		exception: exception.NewException(),
	}
}

func (s *transactionService) Transaction(ctx context.Context, customerId, page, perPage uint) []model.Transaction {
	transactions, err := s.repo.FindByCustomerId(ctx, customerId, page, perPage)
	s.exception.ErrorSkipNotFound(err)
	return transactions
}
