package repository

type UserInterface interface{
 	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) 
 	GetUserByEmail(ctx context.Context, email string) (*domain.User, error) 
}
