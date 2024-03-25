package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"mongofiber/api/routes"
	"mongofiber/database"
	"mongofiber/pkg/article"
	"mongofiber/pkg/item"
	"os"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	database.Connect()
	database.RunMigration()
	articleServiceInstance := article.NewService(article.NewRepo("articles"))
	itemServiceInstance := item.NewService(item.NewRepo("items"))

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.ArticleRouter(api, articleServiceInstance)
	routes.ItemRouter(api, itemServiceInstance)

	log.Fatal(app.Listen(host))
}
