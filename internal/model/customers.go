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
		ID:    c.ID,
		Name:  c.FullName,
		Email: c.Email,
	}
}

type CustomerResponse struct {
	ID           uint      `json:"id"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	LegalName    string    `json:"legal_name"`
	PlaceBirth   string    `json:"place_birth"`
	DateBirth    time.Time `json:"date_birth"`
	Salary       float64   `json:"salary"`
	IdentityFile string    `json:"identity_file"`
	SelfieFile   string    `json:"selfie_file"`
}

func (c Customer) ToResponse() CustomerResponse {
	return CustomerResponse{
		ID:           c.ID,
		NIK:          c.NIK,
		FullName:     c.FullName,
		Email:        c.Email,
		LegalName:    c.LegalName,
		PlaceBirth:   c.PlaceBirth,
		DateBirth:    c.DateBirth,
		Salary:       c.Salary,
		IdentityFile: c.IdentityFile,
		SelfieFile:   c.SelfieFile,
	}
}
