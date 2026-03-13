package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewDatabase(dbUrl string) (*Database, error) {
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	// Test the connection to the database.
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: conn}, nil
}
