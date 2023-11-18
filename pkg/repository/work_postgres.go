package repository

import (
	"database/sql"
	"fmt"
	"strings"

	TonWork "github.com/Glebegor/Ton-Work-Back/structint"
)

type WorkPostgres struct {
	db *sql.DB
}

func NewWorkPostgres(db *sql.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}
func (r *WorkPostgres) GetAll() ([]TonWork.Work, error) {
	var data []TonWork.Work
	query := fmt.Sprintf("SELECT * FROM %s", Table_works)
	rows, err := r.db.Query(query)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var work TonWork.Work
		err := rows.Scan(&work.Id, &work.Title, &work.Description, &work.Text, &work.Tags, &work.Technologies, &work.Company, &work.Price, &work.ExperienceLevel, &work.Type_of_job, &work.Invites, &work.Rating)
		if err != nil {
			return nil, err
		}
		data = append(data, work)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
func (r *WorkPostgres) Create(userId int, data TonWork.Work) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createWorkQuery := fmt.Sprintf("INSERT INTO %s (title, description, text, tags, technologies, company, price, experienceLevel, type_of_job, invites, rating) OUTPUT inserted.id VALUES (@title, @description, @text, @tags, @technologies, @company, @price, @experienceLevel, @type_of_job, @invites, @rating)", Table_works)
	row := tx.QueryRow(createWorkQuery, sql.Named("title", data.Title), sql.Named("description", data.Description), sql.Named("text", data.Text), sql.Named("tags", data.Tags), sql.Named("technologies", data.Technologies), sql.Named("company", data.Company), sql.Named("price", data.Price), sql.Named("experienceLevel", data.ExperienceLevel), sql.Named("type_of_job", data.Type_of_job), sql.Named("invites", data.Invites), sql.Named("rating", data.Rating))
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createWorkUserQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_works) VALUES (@userId, @id)", Table_users_to_works)
	_, err = tx.Exec(createWorkUserQuery, sql.Named("userId", userId), sql.Named("id", id))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *WorkPostgres) GetById(id int) (TonWork.Work, error) {
	var data TonWork.Work
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=@id", Table_works)
	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(&data.Id, &data.Title, &data.Description, &data.Text, &data.Tags, &data.Technologies, &data.Company, &data.Price, &data.ExperienceLevel, &data.Type_of_job, &data.Invites, &data.Rating)
	return data, err
}
func (r *WorkPostgres) Update(id string, input TonWork.WorkUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprint("title=@Title"))
		args = append(args, sql.Named("Title", *input.Title))
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprint("description=@Description"))
		args = append(args, sql.Named("Description", *input.Description))
		argId++
	}
	if input.Text != nil {
		setValues = append(setValues, fmt.Sprint("text=@Text"))
		args = append(args, sql.Named("Text", *input.Text))
		argId++
	}
	if input.Tags != nil {
		setValues = append(setValues, fmt.Sprint("tags=@Tags"))
		args = append(args, sql.Named("Tags", *input.Tags))
		argId++
	}
	if input.Technologies != nil {
		setValues = append(setValues, fmt.Sprint("technologies=@Technologies"))
		args = append(args, sql.Named("Technologies", *input.Technologies))
		argId++
	}
	if input.Company != nil {
		setValues = append(setValues, fmt.Sprint("company=@Company"))
		args = append(args, sql.Named("Company", *input.Company))
		argId++
	}
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprint("price=@Price"))
		args = append(args, sql.Named("Price", *input.Price))
		argId++
	}
	if input.ExperienceLevel != nil {
		setValues = append(setValues, fmt.Sprint("experienceLevel=@ExperienceLevel"))
		args = append(args, sql.Named("ExperienceLevel", *input.ExperienceLevel))
		argId++
	}
	if input.Type_of_job != nil {
		setValues = append(setValues, fmt.Sprint("type_of_job=@Type_of_job"))
		args = append(args, sql.Named("Type_of_job", *input.Type_of_job))
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%s", Table_works, setQuery, id)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *WorkPostgres) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=@id", Table_works)
	_, err := r.db.Exec(query, sql.Named("id", id))
	return err
}
