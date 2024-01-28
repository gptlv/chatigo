package usecase

import (
	"context"
	"fmt"

	"github.com/gptlv/chatigo/server/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) Login(c context.Context, req *domain.User) (*domain.User, error) {
	user, err := uu.userRepo.GetUserByEmail(c, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user by email: %w", err)
	}

	err = CheckPasswordWithHash(req.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to check password with hash: %w", err)
	}

	return &domain.User{Username: user.Username, ID: user.ID}, nil
}

func CheckPasswordWithHash(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
