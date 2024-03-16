package routes

import (
	"github.com/gofiber/fiber/v2"
	"mongofiber/api/handlers"
	"mongofiber/pkg/article"
)

func ArticleRouter(app fiber.Router, service article.Service) {
	app.Get("/articles", handlers.GetArticles(service))
	app.Get("/article", handlers.GetArticle(service))
	app.Post("/article", handlers.AddArticle(service))
	app.Put("/article", handlers.UpdateArticle(service))
	app.Delete("/article", handlers.RemoveArticle(service))
}
