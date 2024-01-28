package postgres

import (
	"database/sql"

	urepository "github.com/gptlv/chatigo/server/internal/interfaces/repository"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) urepository.UserRepository {
	return &repository{db: db}
}
