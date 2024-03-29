package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"mongofiber/api/presenter"
	"mongofiber/pkg/article"
	"mongofiber/pkg/entities"
	"net/http"
)

func GetArticles(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		results, err := service.ReadArticles()
		if err != nil {
			fmt.Println("Error calling service")
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticlesSuccessResponse(results))
	}
}

func GetArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		parsedId, err := uuid.Parse(id)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(errors.New(
				"Please specify id")))
		}

		result, err := service.ReadArticle(parsedId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(result))
	}
}

func AddArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Article

		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}

		if requestBody.Title == "" || requestBody.Description == "" || requestBody.Content == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(errors.New(
				"Please specify title and author")))
		}

		resultArticle, err := service.CreateArticle(&requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(resultArticle))
	}
}

func UpdateArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Article

		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}

		if requestBody.ID.String() == "" || requestBody.Title == "" || requestBody.Description == "" || requestBody.Content == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(errors.New(
				"Please specify title, description and content")))
		}

		result, err := service.UpdateArticle(&requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(result))
	}
}

func RemoveArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		articleId, err := uuid.Parse(id)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(errors.New(
				"Please specify id")))
		}

		err = service.DeleteArticle(articleId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"err":    nil,
		})
	}
}
