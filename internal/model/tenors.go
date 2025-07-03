package model

import "time"

type Tenor struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `gorm:"not null;index" json:"customer_id"`
	Month      uint8     `gorm:"not null" json:"month"`
	Amount     float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Customer Customer `gorm:"foreignKey:CustomerID" json:"-"`
}

func (Tenor) TableName() string {
	return "tenors"
}
