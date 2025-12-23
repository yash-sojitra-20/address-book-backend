package services

import (
	"errors"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
)

type AuthService struct{
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService{
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) error {
	exist, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("email already registered")
	}

	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email: email,
		Password: hashedPass,
	}
	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, user.Email)
}