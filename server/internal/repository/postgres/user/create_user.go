package postgres

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const queryCreateUser = `
INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id
`

func (r *repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var lastInsertId int

	err := r.db.QueryRowContext(ctx, queryCreateUser, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	user.ID = int64(lastInsertId)

	return user, nil
}
