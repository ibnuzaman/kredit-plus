package service

import "kredit-plus/internal/repository"

type authService struct {
	repo repository.AuthRepository
}

type AuthService interface {
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}
