package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mongofiber/pkg/entities"
)

// Article is the presenter object which will be passed in the response by Handler
type Article struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}

// ArticleSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func ArticleSuccessResponse(data *entities.Article) *fiber.Map {
	Article := Article{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Description,
	}
	return &fiber.Map{
		"status": true,
		"data":   Article,
		"error":  nil,
	}
}

// ArticlesSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ArticlesSuccessResponse(data *[]Article) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// ArticleErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ArticleErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
