package service

import (
	"context"
	"fmt"
	"kredit-plus/constant"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
	"kredit-plus/internal/util"
)

type loanService struct {
	repo      repository.LoanRepository
	tenorRepo repository.TenorRepository
	exception exception.Exception
}

type LoanService interface {
	List(ctx context.Context, customerId, page, perPage uint) []model.LoanResponse
	Create(ctx context.Context, customerId uint, req model.CreateLoanRequest)
}

func NewLoanService(repo repository.LoanRepository, tenorRepo repository.TenorRepository) LoanService {
	return &loanService{
		repo:      repo,
		tenorRepo: tenorRepo,
		exception: exception.NewException(),
	}
}

func (s *loanService) List(ctx context.Context, customerId, page, perPage uint) []model.LoanResponse {
	loans, err := s.repo.FindByCustomerId(ctx, customerId, page, perPage)
	s.exception.ErrorSkipNotFound(err)

	res := []model.LoanResponse{}
	for _, loan := range loans {
		res = append(res, loan.ToResponse())
	}
	return res
}

func (s *loanService) Create(ctx context.Context, customerId uint, req model.CreateLoanRequest) {
	last, err := s.repo.GetLastLoanByCustomerId(ctx, customerId)
	s.exception.ErrorSkipNotFound(err)

	if last != nil && last.TenorMonths != uint8(last.TotalPaid) {
		s.exception.UnprocessableEntity("You have an active loan that is not fully paid yet")
	}

	tenors, err := s.tenorRepo.FindByCustomerId(ctx, customerId)
	s.exception.ErrorSkipNotFound(err)
	s.exception.UnprocessableEntityBool(len(tenors) == 0, "No tenor available for this customer")

	mapTenor := make(map[uint8]float64)
	for _, tenor := range tenors {
		mapTenor[tenor.Month] = tenor.Amount
	}

	tenorPerMonthAmount, isOk := mapTenor[req.TenorMonths]
	if !isOk {
		s.exception.UnprocessableEntity(fmt.Sprintf("Tenor %d months is not available for this customer", req.TenorMonths))
	}

	if req.OTR > (tenorPerMonthAmount * float64(req.TenorMonths)) {
		s.exception.UnprocessableEntity(fmt.Sprintf("OTR exceeds the maximum limit. You can set up to %s", util.Rupiah(tenorPerMonthAmount*float64(req.TenorMonths))))
	}

	payload := model.Loan{
		CustomerID:        customerId,
		OTR:               req.OTR,
		TenorMonths:       req.TenorMonths,
		InstallmentAmount: req.OTR * float64(constant.InstallmentFeePercentage),
		AdminFee:          req.OTR * float64(constant.AdminFeePercentage),
		TotalPaid:         0,
		AssetsName:        req.AssetsName,
	}

	err = s.repo.Create(ctx, &payload)
	s.exception.Error(err)
}
