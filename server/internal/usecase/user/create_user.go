package usecase

import (
	"context"
	"fmt"

	"github.com/gptlv/chatigo/server/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) CreateUser(c context.Context, u *domain.User) (*domain.User, error) {
	// проверить полноценность данных, если данные не полные то вернуть ошибку
	// if u == nil {
	// 	return nil, fmt.Errorf(c, "user is nil")
	// }

	// if u.Username == "" {
	// 	return error
	// }

	// if u.Email == "" {
	// 	return error
	// }

	// if u.Password == "" {
	// 	return error
	// }

	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return nil, fmt.Errorf("cann`t create password hash: %w", err)
	}

	user := &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}

	user, err = uu.userRepo.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
