package model

import (
	"kredit-plus/constant"
	"time"
)

type Loan struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID        uint      `gorm:"not null" json:"customer_id"`
	OTR               float64   `gorm:"column:otr;type:decimal(12,2);not null" json:"otr"`
	AdminFee          float64   `gorm:"type:decimal(12,2);not null" json:"admin_fee"`
	InstallmentAmount float64   `gorm:"type:decimal(12,2);not null" json:"installment_amount"`
	AssetsName        string    `gorm:"type:varchar(32);not null" json:"assets_name"`
	TenorMonths       uint8     `gorm:"not null" json:"tenor_months"`
	TotalPaid         int       `gorm:"not null;default:0" json:"total_paid"`
	CreatedAt         time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	Customer     Customer      `gorm:"foreignKey:CustomerID" json:"customer"`
	Transactions []Transaction `gorm:"foreignKey:LoanID" json:"transactions"`
}

func (Loan) TableName() string {
	return "loans"
}

type LoanResponse struct {
	ID                uint      `json:"id"`
	OTR               float64   `json:"otr"`
	AdminFee          float64   `json:"admin_fee"`
	InstallmentAmount float64   `json:"installment_amount"`
	TotalAmount       float64   `json:"total_amount"`
	AssetsName        string    `json:"assets_name"`
	TenorMonths       uint8     `json:"tenor_months"`
	TotalPaid         int       `json:"total_paid"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
}

func (l Loan) ToResponse() LoanResponse {
	total := l.OTR + l.AdminFee + l.InstallmentAmount
	return LoanResponse{
		ID:                l.ID,
		OTR:               l.OTR,
		AdminFee:          l.AdminFee,
		InstallmentAmount: l.InstallmentAmount,
		TotalAmount:       total + (total * constant.InterestPercentage * float64(l.TenorMonths)),
		AssetsName:        l.AssetsName,
		TenorMonths:       l.TenorMonths,
		TotalPaid:         l.TotalPaid,
		StartDate:         l.CreatedAt,
		EndDate:           l.CreatedAt.AddDate(0, int(l.TenorMonths), 0),
	}
}

type CreateLoanRequest struct {
	AssetsName  string  `json:"assets_name" validate:"required" example:"Credit Motorcycle"`
	OTR         float64 `json:"otr" validate:"required,min=0" example:"15000000"`
	TenorMonths uint8   `json:"tenor_months" validate:"required,min=1,max=32" example:"18"`

	AdminFee          float64 `json:"-"`
	InstallmentAmount float64 `json:"-"`
	CustomerID        uint    `json:"-"`
}
