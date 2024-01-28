package usecase

import (
	"context"
	"fmt"

	"github.com/gptlv/chatigo/server/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) CreateUser(c context.Context, u *domain.User) (*domain.User, error) {
	if u == nil {
		return nil, fmt.Errorf("Empty user")
	}

	if u.Username == "" {
		return nil, fmt.Errorf("Empty username")
	}

	if u.Email == "" {
		return nil, fmt.Errorf("Empty email")
	}

	if u.Password == "" {
		return nil, fmt.Errorf("Empty password")
	}

	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a password hash: %w", err)
	}

	user := &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}

	user, err = uu.userRepo.CreateUser(c, user)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a user: %w", err)
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to generate a hash: %w", err)
	}

	return string(hashedPassword), nil
}
