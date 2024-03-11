package routes

import (
	"github.com/gofiber/fiber/v2"
	"mongofiber/api/handlers"
	"mongofiber/pkg/item"
)

// ItemRouter is the Router for GoFiber App
func ItemRouter(app fiber.Router, service item.Service) {
	app.Get("/items", handlers.GetItems(service))
	app.Post("/items", handlers.AddItem(service))
	app.Put("/items", handlers.UpdateItem(service))
	app.Delete("/items", handlers.RemoveItem(service))
}
