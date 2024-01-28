package postgres

import (
	"context"
	"fmt"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const queryGetUserByEmail = `
SELECT id, email, username, password FROM users WHERE email = $1
`

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := domain.User{}

	err := r.db.QueryRowContext(ctx, queryGetUserByEmail, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute a row query: %w", err)
	}

	return &user, nil
}
