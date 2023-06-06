package repository

import (
	"fmt"

	TonWork "github.com/TonWork/back"
	"github.com/jmoiron/sqlx"
)

type WorkPostgres struct {
	db *sqlx.DB
}

func NewWorkPostgres(db *sqlx.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}
func (r *WorkPostgres) GetAll() ([]TonWork.Work, error) {
	var data []TonWork.Work
	query := fmt.Sprintf("SELECT * FROM %s", Table_works)

	err := r.db.Select(&data, query)
	return data, err
}
func (r *WorkPostgres) Create(int, TonWork.Work) error {

}
