package model

import (
	"time"
)

type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	NIK          string    `gorm:"type:varchar(16);not null;unique"`
	FullName     string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(128);not null;unique"`
	Password     string    `gorm:"type:varchar(72);not null"`
	LegalName    string    `gorm:"type:varchar(100);not null"`
	PlaceBirth   string    `gorm:"type:varchar(32);not null"`
	DateBirth    time.Time `gorm:"type:date;not null"`
	Sallary      float64   `gorm:"type:decimal(12,2);not null"`
	IdentityFile string    `gorm:"type:varchar(64);not null"`
	SelfieFile   string    `gorm:"type:varchar(64);not null"`
	CreatedAt    time.Time `gorm:"type:date;not null"`
	UpdatedAt    time.Time `gorm:"type:date"`
}

func (Customer) TableName() string {
	return "customers"
}
