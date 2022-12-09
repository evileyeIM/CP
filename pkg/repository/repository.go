package repository

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type Note interface {
	CreateNote(userId int, name string, date string, value string) (models.Note, error)
	DeleteNote(userId int, name string) (models.Note, error)
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          NewNotePostgres(db),
	}
}
