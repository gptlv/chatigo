package restapi

import (
	uusecase "github.com/gptlv/chatigo/server/internal/interfaces/usecase"
)

type UserHandler struct {
	userUsecase uusecase.UserUsecase
}

func NewUserHandler(uu uusecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}
