package user

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gptlv/chatigo/server/util"
)

type UseCase struct{
	rep repository.UserInterface
}

func NewUseCase(rep repository.UserInterface) *UseCase{
	return *UseCase{
		rep:	rep,
	}
}
