package models

import "github.com/google/uuid"

type Book struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Year   int       `json:"year"`
	Author string    `json:"author"`
}
