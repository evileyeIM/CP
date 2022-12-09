package repository

import (
	"fmt"
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/jmoiron/sqlx"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) CreateNote(userId int, noteName string, noteDate string, noteValue string) (models.Note, error) {
	var note models.Note
	query := fmt.Sprintf("INSERT INTO %s (userId, noteName, noteDate, noteValue)", notesTable)
	err := r.db.Get(&note, query)
	return note, err
}

func (r *NotePostgres) DeleteNote(userId int, nameToDelete string) (models.Note, error) {
	var note models.Note
	query := fmt.Sprintf("DELETE FROM %s (userId, nameToDelete) WHERE userId='userId' AND name='nameToDelete')", notesTable)
	err := r.db.Get(&note, query)
	return note, err
}
