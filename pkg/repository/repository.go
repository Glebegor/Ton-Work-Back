package repository

import (
	sqlx "github.com/jmoiron/sqlx"
)

type Authorization interface {
}
type Work interface {
}
type Posts interface {
}
type Subscribes interface {
}
type Repository struct {
	Authorization
	Work
	Posts
	Subscribes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
