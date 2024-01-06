package user

func (uc *UseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// проверить полноценность данных, если данные не полные то вернуть ошибку
	if req == nil{
		return nil, fmt.Errorf(ctx, "user is nil")
	}

	if req.Username == ""{
		return error
	}

	if req.Email == ""{
		return error
	}

	if req.Password == ""{
		return error
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf(ctx, "cann`t create password hash: %w", err)
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err = uc.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}


func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}