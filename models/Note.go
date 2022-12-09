package models

type Note struct {
	Id        int    `json:"id" db:"id"`
	UserId    string `json:"userId" db:"userId" binding:"required"`
	Name      string `json:"name" db:"name" binding:"required"`
	NoteValue string `json:"noteValue" db:"noteValue" binding:"required"`
	Date      string `json:"date" db:"date" binding:"required"`
}
