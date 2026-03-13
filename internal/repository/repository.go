package repository

import "database/sql"

type Repository struct {
	db *sql.DB
	*UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db, UserRepository: &UserRepository{db: db}}
}
