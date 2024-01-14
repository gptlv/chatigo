package restapi

import (
	"github.com/gptlv/chatigo/server/internal/domain"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}
