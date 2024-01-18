package postgres

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const getAllUsersQuery = `
SELECT id FROM users
`

func (ur *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := ur.db.QueryContext(ctx, getAllUsersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		rows.Scan(&user.Id, &user.Email)
		users = append(users, user)
	}

	return users, nil
}
