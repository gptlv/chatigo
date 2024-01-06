package user

import (
	"context"
	"database/sql"
)

const queryGetUserByEmail = `
SELECT id, email, username, password FROM users WHERE email = $1
`

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}