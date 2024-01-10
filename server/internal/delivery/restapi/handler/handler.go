package handler

import (
	"github.com/gptlv/chatigo/server/internal/domain"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewHandler(uu domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}
