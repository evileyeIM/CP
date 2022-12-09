package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type inputCreateNote struct {
	UserId int    `json:"usedId"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Value  string `json:"value"`
}

type inputDeleteNote struct {
	UserId int    `json:"usedId"`
	Name   string `json:"name"`
}

func (h *Handler) createNote(c *gin.Context) {
	fmt.Println("createNote")

	var input inputCreateNote

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	note, err := h.services.Note.CreateNote(input.UserId, input.Name, input.Date, input.Value)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"userId": note.UserId, "name": note.Name, "date": note.Date, "value": note.NoteValue})
}

func (h *Handler) deleteNote(c *gin.Context) {
	fmt.Println("deleteNote")

	var input inputDeleteNote

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	note, err := h.services.Note.DeleteNote(input.UserId, input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"userId": note.UserId, "name": note.Name})
}
