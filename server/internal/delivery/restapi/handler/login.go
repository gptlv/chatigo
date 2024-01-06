package handler 

func (h *Handler) Login(c *gin.Context) {
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUseCase.Login(c.Request.Context(), &domain.User{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil{
		error
	}

	c.SetCookie("jwt", ss, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &LoginUserRes{
		Username: user.Username, 
		ID: user.ID,
	})
}

type LoginUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
}