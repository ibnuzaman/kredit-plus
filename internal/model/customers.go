package model

import (
	"time"
)

type Customer struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NIK          string    `gorm:"type:varchar(16);not null;uniqueIndex" json:"nik"`
	FullName     string    `gorm:"type:varchar(100);not null" json:"full_name"`
	Email        string    `gorm:"type:varchar(128);not null;uniqueIndex" json:"email"`
	Password     string    `gorm:"type:varchar(72);not null" json:"-"`
	LegalName    string    `gorm:"type:varchar(100);not null" json:"legal_name"`
	PlaceBirth   string    `gorm:"type:varchar(32);not null" json:"place_birth"`
	DateBirth    time.Time `gorm:"type:date;not null" json:"date_birth"`
	Salary       float64   `gorm:"type:decimal(12,2);not null" json:"salary"`
	IdentityFile string    `gorm:"type:varchar(64);not null" json:"identity_file"`
	SelfieFile   string    `gorm:"type:varchar(64);not null" json:"selfie_file"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Tenors       []Tenor       `gorm:"foreignKey:CustomerID" json:"tenors,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:CustomerID" json:"transactions,omitempty"`
}

func (Customer) TableName() string {
	return "customers"
}

func (c Customer) ToAuthMe() AuthMe {
	return AuthMe{
		ID:         c.ID,
		NIK:        c.NIK,
		FullName:   c.FullName,
		LegalName:  c.LegalName,
		Email:      c.Email,
		Salary:     c.Salary,
		SelfieFile: c.SelfieFile,
	}
}
