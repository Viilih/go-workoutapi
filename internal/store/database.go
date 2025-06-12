package store

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenDatabase() (*sql.DB, error) {
	db, error := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres sslmode=disable")

	if error != nil {
		return nil, fmt.Errorf("error opening database: %w", error)
	}
	fmt.Println("Connected to database")
	return db, nil
}
