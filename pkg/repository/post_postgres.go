package repository

import (
	"fmt"

	TonWork "github.com/TonWork/back"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}
func (r *PostPostgres) GetAll() ([]TonWork.Post, error) {
	var data []TonWork.Post
	query := fmt.Sprintf("SELECT * FROM %s", Table_posts)

	err := r.db.Select(&data, query)
	return data, err
}
func (r *PostPostgres) Create(userId int, data TonWork.Post) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description, text, tags, rating) VALUES ($1, $2, $3, $4, $5) RETURNING id", Table_works)
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
