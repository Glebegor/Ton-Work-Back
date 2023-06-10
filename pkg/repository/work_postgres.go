package repository

import (
	"fmt"
	"strings"

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
func (r *WorkPostgres) GetById(id int) (TonWork.Work, error) {
	var data TonWork.Work
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", Table_works)
	err := r.db.Get(&data, query, id)
	return data, err
}
func (r *WorkPostgres) Update(id string, input TonWork.WorkUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}
	if input.Tags != nil {
		setValues = append(setValues, fmt.Sprintf("tags=$%d", argId))
		args = append(args, *input.Tags)
		argId++
	}
	if input.Technologies != nil {
		setValues = append(setValues, fmt.Sprintf("technologies=$%d", argId))
		args = append(args, *input.Technologies)
		argId++
	}
	if input.Company != nil {
		setValues = append(setValues, fmt.Sprintf("company=$%d", argId))
		args = append(args, *input.Company)
		argId++
	}
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	if input.ExperienceLevel != nil {
		setValues = append(setValues, fmt.Sprintf("experienceLevel=$%d", argId))
		args = append(args, *input.ExperienceLevel)
		argId++
	}
	if input.Type_of_job != nil {
		setValues = append(setValues, fmt.Sprintf("type_of_job=$%d", argId))
		args = append(args, *input.Type_of_job)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%s", Table_works, setQuery, id)
	_, err := r.db.Exec(query, args...)
	return err
}
