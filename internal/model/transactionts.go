package model

import "time"

type Transaction struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID         uint      `gorm:"not null" json:"loan_id"`
	Amount         float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	InterestAmount float64   `gorm:"type:decimal(12,2);not null" json:"interest_amount"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:datetime" json:"updated_at"`

	Loan Loan `gorm:"foreignKey:LoanID" json:"loan,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}
