package user

func (uc *UseCase) Login(c context.Context, req *domain.User) (*domain.User, error) {
	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Error(ctx, "err with getting email: %w", err)
	}

	err = CheckPasswordWithHash(req.Password, u.Password)
	if err != nil {
		return nil, fmt.Error(ctx, "err with check password with hash: %w", err)
	}

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

func CheckPasswordWithHash(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


/*
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
*/