package entities

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EntityId    uuid.UUID `json:"entityId"`
	EntityType  string    `json:"entityType"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
