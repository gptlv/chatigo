package handler

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUseCase user.UseCase
}

func NewHandler(userUseCase user.UseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
	}
}
