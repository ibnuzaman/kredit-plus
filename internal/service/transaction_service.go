package service

import (
	"context"
	"fmt"
	"kredit-plus/constant"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
	"kredit-plus/internal/util"
	"math"
)

type transactionService struct {
	repo      repository.TransactionRepository
	loanRepo  repository.LoanRepository
	exception exception.Exception
}

type TransactionService interface {
	List(ctx context.Context, customerId, page, perPage uint) []model.TransactionResponse
	Create(ctx context.Context, customerId uint, req model.CreateTransactionRequest)
}

func NewTransactionService(repo repository.TransactionRepository, loanRepo repository.LoanRepository) TransactionService {
	return &transactionService{
		repo:      repo,
		loanRepo:  loanRepo,
		exception: exception.NewException(),
	}
}

func (s *transactionService) List(ctx context.Context, customerId, page, perPage uint) []model.TransactionResponse {
	transactions, err := s.repo.FindByCustomerId(ctx, customerId, page, perPage)
	s.exception.ErrorSkipNotFound(err)

	res := []model.TransactionResponse{}
	for _, transaction := range transactions {
		res = append(res, transaction.ToResponse())
	}
	return res
}

func (s *transactionService) Create(ctx context.Context, customerId uint, req model.CreateTransactionRequest) {
	loan, err := s.loanRepo.GetById(ctx, req.LoanID)
	s.exception.ErrorSkipNotFound(err)
	s.exception.UnprocessableEntityBool(loan == nil, "Loan not found")

	if loan.TenorMonths == uint8(loan.TotalPaid) {
		s.exception.UnprocessableEntity("You have fully paid this loan")
	}

	totalAmount := loan.OTR + loan.InstallmentAmount + loan.AdminFee
	interestAmount := totalAmount * float64(constant.InterestPercentage)
	totalAmountPerMonth := totalAmount / float64(loan.TenorMonths)

	shouldBePay := math.Round(totalAmountPerMonth + interestAmount)
	if req.Amount != shouldBePay {
		s.exception.UnprocessableEntity(fmt.Sprintf("Transaction amount does not match the total amount of the loan. You should pay %s", util.Rupiah(shouldBePay)))
	}

	payload := model.Transaction{
		CustomerID:     customerId,
		LoanID:         loan.ID,
		Amount:         totalAmountPerMonth,
		InterestAmount: interestAmount,
	}

	err = s.repo.Create(ctx, &payload)
	s.exception.Error(err)
}
