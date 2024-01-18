package postgres

import (
	"context"

	"github.com/gptlv/chatigo/server/internal/domain"
)

const deleteUserQuery = `
DELETE FROM users WHERE id = $1
`

func (ur *userRepository) DeleteUser(ctx context.Context, user *domain.User) error {

	err := ur.db.QueryRowContext(ctx, deleteUserQuery, user.Id).Scan()
	if err != nil {
		return err
	}

	return nil
}
