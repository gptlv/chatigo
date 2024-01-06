package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func NewDatabase() (sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:root@localhost:5433/chatigo?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}