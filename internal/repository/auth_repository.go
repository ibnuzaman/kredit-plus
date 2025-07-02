package repository

import "database/sql"

type authRepository struct {
	db *sql.DB
}

type AuthRepository interface {
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
