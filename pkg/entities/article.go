package entities

import "github.com/google/uuid"

type Article struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}

type Articles struct {
	Employees []Article `json:"article"`
}
