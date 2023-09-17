package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
