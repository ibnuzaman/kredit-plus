package model

import "time"

type Tenor struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `gorm:"not null;index" json:"customer_id"`
	Month      uint8     `gorm:"not null" json:"month"`
	Amount     float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

func (Tenor) TableName() string {
	return "tenors"
}

type TenorResponse struct {
	ID     uint    `json:"id"`
	Month  uint8   `json:"month"`
	Amount float64 `json:"amount"`
}

func (t Tenor) ToResponse() TenorResponse {
	return TenorResponse{
		ID:     t.ID,
		Month:  t.Month,
		Amount: t.Amount,
	}
}
