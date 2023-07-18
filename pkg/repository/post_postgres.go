package repository

import (
	"database/sql"
	"fmt"
	"strings"

	TonWork "github.com/TonWork/back"
)

type PostPostgres struct {
	db *sql.DB
}

func NewPostPostgres(db *sql.DB) *PostPostgres {
	return &PostPostgres{db: db}
}
func (r *PostPostgres) GetAll() ([]TonWork.Post, error) {
	var data []TonWork.Post
	query := fmt.Sprintf("SELECT * FROM %s", Table_posts)

	err := r.db.QueryRow(query).Scan(&data)
	return data, err
}
func (r *PostPostgres) Create(userId int, data TonWork.Post) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description, text, tags, rating) VALUES ($1, $2, $3, $4, $5) RETURNING id", Table_posts)
	row := tx.QueryRow(createPostQuery, data.Title, data.Description, data.Text, data.Tags, data.Rating)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createPostUserQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_posts) VALUES ($1, $2)", Table_users_to_posts)
	_, err = tx.Exec(createPostUserQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *PostPostgres) GetById(id int) (TonWork.Post, error) {
	var data TonWork.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", Table_posts)
	err := r.db.QueryRow(query, id).Scan(&data)
	return data, err
}
func (r *PostPostgres) Update(id string, input TonWork.PostUpdate) error {
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
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%s", Table_posts, setQuery, id)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *PostPostgres) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", Table_posts)
	_, err := r.db.Exec(query, id)
	return err
}
