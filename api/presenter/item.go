package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mongofiber/pkg/entities"
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

func ItemSuccessResponse(data *entities.Item) *fiber.Map {
	Item := Item{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		EntityType:  data.EntityType,
		EntityId:    data.EntityId,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   Item,
		"error":  nil,
	}
}

// ItemsSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ItemsSuccessResponse(data *[]Item) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ItemErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ItemErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
