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
func (r *WorkPostgres) Create(userId int, data TonWork.Work) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createWorkQuery := fmt.Sprintf("INSERT INTO %s (title, description, text, tags, technologies, company, price, experienceLevel, type_of_job, invites, rating) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", Table_works)
	row := tx.QueryRow(createWorkQuery, data.Title, data.Description, data.Text, data.Tags, data.Technologies, data.Company, data.Price, data.ExperienceLevel, data.Type_of_job, data.Invites, data.Rating)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createWorkUserQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_works) VALUES ($1, $2)", Table_users_to_works)
	_, err = tx.Exec(createWorkUserQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
