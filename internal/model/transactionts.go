package model

import "time"

type Transaction struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	CustomerID        uint      `gorm:"not null;index" json:"customer_id"`
	ContractNumber    string    `gorm:"type:varchar(32);not null;index" json:"contract_number"`
	OTR               float64   `gorm:"type:decimal(12,2);not null" json:"otr"`
	AdminFee          float64   `gorm:"type:decimal(12,2);not null" json:"admin_fee"`
	InstallmentAmount float64   `gorm:"type:decimal(12,2);not null" json:"installment_amount"`
	InterestAmount    float64   `gorm:"type:decimal(12,2);not null" json:"interest_amount"`
	AssetsName        string    `gorm:"type:varchar(32);not null" json:"assets_name"`
	TenorMonths       uint8     `gorm:"not null" json:"tenor_months"`
	CreatedAt         time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	Customer Customer `gorm:"foreignKey:CustomerID" json:"-"`
}

func (Transaction) TableName() string {
	return "transactions"
}
