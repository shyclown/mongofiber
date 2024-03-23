package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
	"mongofiber/pkg/item"
	"net/http"
)

func GetItems(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		results, err := service.FetchItems()
		if err != nil {
			fmt.Println("Error calling service")
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemsSuccessResponse(results))
	}
}

func GetItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		parsedId, err := uuid.Parse(id)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(errors.New(
				"Please specify id")))
		}

		result, err := service.FetchItem(parsedId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemSuccessResponse(result))
	}
}

func AddItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Item

		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(err))
		}

		if requestBody.Title == "" || requestBody.Description == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(errors.New(
				"Please specify title and description")))
		}

		result, err := service.InsertItem(&requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemSuccessResponse(result))
	}
}

func UpdateItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Item

		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(err))
		}

		if requestBody.Id.String() == "" || requestBody.Title == "" || requestBody.Description == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(errors.New(
				"Please specify title, description and content")))
		}

		result, err := service.UpdateItem(&requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemSuccessResponse(result))
	}
}

func RemoveItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		itemId, err := uuid.Parse(id)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(errors.New(
				"Please specify id")))
		}

		err = service.RemoveItem(itemId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"err":    nil,
		})
	}
}
