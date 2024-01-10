package postgres

import (
	"database/sql"

	"github.com/gptlv/chatigo/server/internal/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.UserRepository {
	return &repository{db: db}
}
