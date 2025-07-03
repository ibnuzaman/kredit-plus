package service

import (
	"context"
	"errors"
	"fmt"
	"kredit-plus/config"
	"kredit-plus/exception"
	"kredit-plus/internal/model"
	"kredit-plus/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	repo      repository.AuthRepository
	exception exception.Exception
	conf      *config.Config
}

type AuthService interface {
	Login(ctx context.Context, req model.LoginRequest) model.LoginResponse
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo:      repo,
		exception: exception.NewException(),
		conf:      config.Get(),
	}
}

func (s *authService) Login(ctx context.Context, req model.LoginRequest) model.LoginResponse {
	customer, err := s.repo.FindByEmail(ctx, req.Email)
	const invalidEmailOrPassword = "Email or password is incorrect"
	s.exception.UnauthorizedBool(err != nil && errors.Is(err, gorm.ErrRecordNotFound), invalidEmailOrPassword)
	s.exception.UnauthorizedBool(s.compareHash(customer.Password, req.Password), invalidEmailOrPassword)

	token, err := s.generateToken(customer.ID)
	s.exception.Error(err)

	return model.LoginResponse{
		Token: token,
	}
}

func (s *authService) compareHash(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil
}

func (s *authService) generateToken(customerId uint) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		ID:        fmt.Sprint(customerId),
		ExpiresAt: jwt.NewNumericDate(now.Add(s.conf.Auth.ExpiredDuration)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(s.conf.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
