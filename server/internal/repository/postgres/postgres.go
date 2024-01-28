package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDatabase(source string) (*sql.DB, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("Failed to open a database: %w", err)
	}

	return db, nil
}
