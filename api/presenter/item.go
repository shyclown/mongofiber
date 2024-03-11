package presenter

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongofiber/pkg/entities"
)

// Item is the presenter object which will be passed in the response by Handler
type Item struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
}

// ItemSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func ItemSuccessResponse(data *entities.Item) *fiber.Map {
	Item := Item{
		ID:     data.ID,
		Title:  data.Title,
		Author: data.Author,
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
