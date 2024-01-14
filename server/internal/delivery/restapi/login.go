package restapi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gptlv/chatigo/server/internal/domain"
)

const secretKey = "secret"

type LoginUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (uh *UserHandler) Login(c *gin.Context) {
	var userReq LoginUserReq

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	loggedUser, err := uh.userUsecase.Login(c.Request.Context(), &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(loggedUser.ID)),
		Username: loggedUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(loggedUser.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))

	c.SetCookie("jwt", ss, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{Username: loggedUser.Username, ID: strconv.Itoa(int(loggedUser.ID))}

	c.JSON(http.StatusOK, res)

}
