package service

import (
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUser(email, password string) (int, error)
}

type Note interface {
	CreateNote(userId int, name string, date string, value string) (models.Note, error)
	DeleteNote(userId int, name string) (models.Note, error)
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNote(repos.Note),
	}
}
