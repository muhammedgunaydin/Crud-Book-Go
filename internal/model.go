package internal

import "github.com/google/uuid"

type Book struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Page   int       `json:"page,omitempty"`
	Author string    `json:"author,omitempty"`
}
