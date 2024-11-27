package application

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

type Application struct {
	ID     int    `sql:"id" json:"id"`
	UserId int    `sql:"user_id" json:"userId"`
	Title  string `sql:"title" json:"title"`
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindApplicationById(id int) (*Application, error) {
	application := &Application{}
	query := "SELECT * FROM applications WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&application.ID, &application.UserId, &application.Title)
	if err != nil {
		fmt.Println("error with the query", err)
		return nil, err
	}
	return application, nil
}
