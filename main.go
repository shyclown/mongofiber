package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"mongofiber/api/routes"
	"mongofiber/database"
	"mongofiber/pkg/article"
)

func main() {

	database.Connect()

	articleTable := "articles"
	articleRepo := article.NewRepo(articleTable)
	articleService := article.NewService(articleRepo)

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.ArticleRouter(api, articleService)
	//defer cancel()
	log.Fatal(app.Listen(":8080"))
}
