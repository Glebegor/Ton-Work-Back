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

	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var post TonWork.Post
		err := row.Scan(&post.Id, &post.Title, &post.Description, &post.Text, &post.Tags, &post.Rating)
		if err != nil {
			return nil, err
		}
		data = append(data, post)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return data, err
}
func (r *PostPostgres) Create(userId int, data TonWork.Post) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description, text, tags, rating) OUTPUT inserted.id VALUES (@Title, @Description, @Text, @Tags, @Rationg)", Table_posts)
	row := tx.QueryRow(createPostQuery, sql.Named("Title", data.Title), sql.Named("Description", data.Description), sql.Named("Text", data.Text), sql.Named("Tags", data.Tags), sql.Named("Rationg", data.Rating))
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createPostUserQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_posts) VALUES (@UserId, @Id)", Table_users_to_posts)
	_, err = tx.Exec(createPostUserQuery, sql.Named("UserId", userId), sql.Named("Id", id))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *PostPostgres) GetById(id int) (TonWork.Post, error) {
	var data TonWork.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=@Id", Table_posts)
	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(&data.Id, &data.Title, &data.Description, &data.Text, &data.Tags, &data.Rating)
	return data, err
}
func (r *PostPostgres) Update(id string, input TonWork.PostUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, "title=@Title")
		args = append(args, sql.Named("Title", *input.Title))
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, "description=@Description")
		args = append(args, sql.Named("Description", *input.Description))
		argId++
	}
	if input.Text != nil {
		setValues = append(setValues, "text=@Text")
		args = append(args, sql.Named("Text", *input.Text))
		argId++
	}
	if input.Tags != nil {
		setValues = append(setValues, "tags=@Tags")
		args = append(args, sql.Named("Tags", *input.Tags))
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%s", Table_posts, setQuery, id)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *PostPostgres) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=@Id", Table_posts)
	_, err := r.db.Exec(query, sql.Named("Id", id))
	return err
}
