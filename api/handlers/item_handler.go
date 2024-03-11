package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
	"mongofiber/pkg/item"
	"net/http"
)

// AddItem is handler/controller which creates Items in the ItemShop
func AddItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Item
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		if requestBody.Author == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(errors.New(
				"Please specify title and author")))
		}
		result, err := service.InsertItem(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemSuccessResponse(result))
	}
}

// UpdateItem is handler/controller which updates data of Items in the ItemShop
func UpdateItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Item
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		result, err := service.UpdateItem(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemSuccessResponse(result))
	}
}

// RemoveItem is handler/controller which removes Items from the ItemShop
func RemoveItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		itemID := requestBody.ID
		err = service.RemoveItem(itemID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

// GetItems is handler/controller which lists all Items from the ItemShop
func GetItems(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchItems()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ItemErrorResponse(err))
		}
		return c.JSON(presenter.ItemsSuccessResponse(fetched))
	}
}
