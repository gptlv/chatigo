package usecase

import (
	urepository "github.com/gptlv/chatigo/server/internal/interfaces/repository"
	uusecase "github.com/gptlv/chatigo/server/internal/interfaces/usecase"
)

type userUsecase struct {
	userRepo urepository.UserRepository
}

func NewUserUsecase(ur urepository.UserRepository) uusecase.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}
