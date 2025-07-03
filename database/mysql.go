package database

import (
	"context"
	"database/sql"
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
			logs.Info().Msg("Auto migrating database schemas")
			autoMigrate(logs)
			logs.Info().Msg("MySQL database initialized successfully")
		}
	})
}

func autoMigrate(log *zerolog.Logger) {
	log.Info().Msg("Auto migrating database schemas")
	err := dbGorm.AutoMigrate(model.Customer{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to auto migrate database schemas")
		return
	}
	log.Info().Msg("Database schemas migrated successfully")
}

func Get() *sql.DB {
	return db
}

func GetGorm() *gorm.DB {
	return dbGorm
}
