package service

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNote(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(userId int, name string, date string, value string) (models.Note, error) {
	note, err := s.repo.CreateNote(userId, name, date, value)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (s *NoteService) DeleteNote(userId int, name string) (models.Note, error) {
	note, err := s.repo.DeleteNote(userId, name)
	if err != nil {
		return note, err
	}
	return note, nil
}
