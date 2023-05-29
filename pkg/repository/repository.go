package repository

import (
	sqlx "github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}
