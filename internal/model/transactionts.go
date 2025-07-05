package model

import "time"

type Transaction struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID         uint      `gorm:"not null" json:"loan_id"`
	CustomerID     uint      `gorm:"not null" json:"customer_id"`
	Amount         float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	InterestAmount float64   `gorm:"type:decimal(12,2);not null" json:"interest_amount"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:datetime" json:"updated_at"`

	Loan     Loan     `gorm:"foreignKey:LoanID" json:"loan,omitempty"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type TransactionResponse struct {
	ID             uint      `json:"id"`
	LoanID         uint      `json:"loan_id"`
	Amount         float64   `json:"amount"`
	InterestAmount float64   `json:"interest_amount"`
	Date           time.Time `json:"date"`
}

func (t Transaction) ToResponse() TransactionResponse {
	return TransactionResponse{
		ID:             t.ID,
		LoanID:         t.LoanID,
		Amount:         t.Amount,
		InterestAmount: t.InterestAmount,
		Date:           t.CreatedAt,
	}
}

type CreateTransactionRequest struct {
	LoanID uint    `json:"loan_id" binding:"required" example:"1"`
	Amount float64 `json:"amount" binding:"required,min=0" example:"1000.00"`

	CustomerID uint `json:"-"`
}
