package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kredit-plus/config"
	"kredit-plus/internal/model"
	"kredit-plus/logger"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var onceMySql sync.Once
var db *sql.DB
var dbGorm *gorm.DB

func Init(ctx context.Context, conf *config.Config) {
	onceMySql.Do(func() {
		logs := logger.Get("mysql")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DB.User,
			conf.DB.Password,
			conf.DB.Host,
			conf.DB.Port,
			conf.DB.Name,
		)

		var dbLogger gormLogger.Interface
		if conf.Env.IsDevelopment() {
			dbLogger = gormLogger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gormLogger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  gormLogger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  false,
				},
			)
		}

		config := &gorm.Config{
			Logger: dbLogger,
		}

		var err error
		dbMySql, err := gorm.Open(mysql.Open(dsn), config)
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to connect to MySQL database")
		}

		sqlDB, err := dbMySql.DB()
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to get sql.DB object from GORM")
		}

		sqlDB.SetMaxIdleConns(conf.DB.MaxIdle)
		sqlDB.SetMaxOpenConns(conf.DB.MaxOpen)
		sqlDB.SetConnMaxLifetime(conf.DB.ConnectionLifetime)
		sqlDB.SetConnMaxIdleTime(conf.DB.ConnectionIdle)

		logs.Info().Msg("MySQL database connected")

		db = new(sql.DB)
		db = sqlDB

		dbGorm = new(gorm.DB)
		dbGorm = dbMySql

		if conf.DB.AutoMigrate {
			autoMigrate(logs)
			autoSeeder(logs)
		}
	})
}

func autoMigrate(log *zerolog.Logger) {
	log.Info().Msg("Auto migrating database schemas")
	err := dbGorm.AutoMigrate(&model.Customer{}, &model.Tenor{}, &model.Loan{}, &model.Transaction{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to auto migrate database schemas")
		return
	}
	log.Info().Msg("Database schemas migrated successfully")
}

func autoSeeder(log *zerolog.Logger) {
	log.Info().Msg("Auto seeding database")

	type CustomerSeedData struct {
		Customer model.Customer
		Tenors   []model.Tenor
	}

	seeders := []CustomerSeedData{
		{
			Customer: model.Customer{
				NIK:          "3510010101900001",
				FullName:     "Budi Santoso",
				Email:        "budi@gmail.com",
				Password:     "$2a$12$8rHI4GvSgCr8LDn9c6Y/kO/71biGspofK8Zeyh/jDZsevMJRnB/Jy", // password
				LegalName:    "Budi Santoso",
				PlaceBirth:   "Jakarta",
				DateBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				Salary:       5000000,
				IdentityFile: "9c35760e-4d3e-45af-9fdc-f4fdd9cf3185.jpg",
				SelfieFile:   "ca7022c1-685d-4903-94f6-1f469c135f72.jpg",
			},
			Tenors: []model.Tenor{
				{Month: 1, Amount: 100000},
				{Month: 2, Amount: 200000},
				{Month: 3, Amount: 500000},
				{Month: 6, Amount: 700000},
			},
		},
		{
			Customer: model.Customer{
				NIK:          "3510010202920002",
				FullName:     "Annisa Fitriani",
				Email:        "annisa@gmail.com",
				Password:     "$2a$12$8rHI4GvSgCr8LDn9c6Y/kO/71biGspofK8Zeyh/jDZsevMJRnB/Jy", // password
				LegalName:    "Annisa Fitriani",
				PlaceBirth:   "Surabaya",
				DateBirth:    time.Date(1992, 2, 2, 0, 0, 0, 0, time.UTC),
				Salary:       15000000,
				IdentityFile: "f19efd94-484d-4507-b961-19c8d354d7ef.jpg",
				SelfieFile:   "220ab59f-b476-4064-9227-234e36842284.jpg",
			},
			Tenors: []model.Tenor{
				{Month: 1, Amount: 1000000},
				{Month: 2, Amount: 1200000},
				{Month: 3, Amount: 1500000},
				{Month: 6, Amount: 2000000},
			},
		},
	}

	tx := dbGorm.Begin()
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("Failed to start transaction for auto seeding")
		return
	}

	for _, data := range seeders {
		var existingCustomer model.Customer
		if err := tx.Where("email = ?", data.Customer.Email).First(&existingCustomer).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Create(&data.Customer).Error; err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Failed to create customer during auto seeding")
				return
			}
		} else if err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("Failed to check existing customer during auto seeding")
			return
		} else {
			data.Customer.ID = existingCustomer.ID
		}

		for _, tenor := range data.Tenors {
			tenor.CustomerID = data.Customer.ID
			var existingTenor model.Tenor
			if err := tx.Where("customer_id = ? AND month = ?", tenor.CustomerID, tenor.Month).First(&existingTenor).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&tenor).Error; err != nil {
					tx.Rollback()
					log.Error().Err(err).Msg("Failed to create tenor during auto seeding")
					return
				}
			} else if err != nil {
				tx.Rollback()
				log.Error().Err(err).Msg("Failed to check existing tenor during auto seeding")
				return
			}
		}
	}

	err := tx.Commit().Error
	if err != nil {
		log.Error().Err(err).Msg("Failed to commit transaction for auto seeding")
		return
	}
}

func Get() *sql.DB {
	return db
}

func GetGorm() *gorm.DB {
	return dbGorm
}
