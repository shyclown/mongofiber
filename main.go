package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"mongofiber/api/routes"
	"mongofiber/database"
	"mongofiber/pkg/article"
	"mongofiber/pkg/item"
)

func main() {

	// load .env file
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}
	//fmt.Println("Environment loaded: ", os.Getenv("ENVIRONMENT"))

	database.Connect()
	database.RunMigration()

	articleTable := "articles"
	articleRepo := article.NewRepo(articleTable)
	articleService := article.NewService(articleRepo)

	itemTable := "items"
	itemRepo := item.NewRepo(itemTable)
	itemService := item.NewService(itemRepo)

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.ArticleRouter(api, articleService)
	routes.ItemRouter(api, itemService)
	//defer cancel()
	// log.Fatal(app.Listen("0.0.0.0:8080"))
	log.Fatal(app.Listen("localhost:8080"))
}
