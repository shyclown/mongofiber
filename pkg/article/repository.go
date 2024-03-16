package article

import (
	"fmt"
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/database"
	"mongofiber/pkg/entities"
)

type Repository interface {
	InsertArticle(article *entities.Article) (*entities.Article, error)
	FetchArticles() (*[]presenter.Article, error)
	UpdateArticle(article *entities.Article) (*entities.Article, error)
	DeleteArticle(ID uuid.UUID) error
}

type repository struct {
	Table string
}

var table string = "articles"

// NewRepo is the single instance repo that is being created.
func NewRepo(table string) Repository {
	return &repository{
		Table: table,
	}
}

func (r *repository) InsertArticle(article *entities.Article) (*entities.Article, error) {
	article.ID = uuid.New()

	// Call sql here
	_, err := database.DB.Query(
		"INSERT INTO articles (ID, TITLE, DESCRIPTION,CONTENT) VALUES (?, ?, ?, ?)",
		article.ID, article.Title, article.Description, article.Content,
	)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *repository) FetchArticles() (*[]presenter.Article, error) {
	fmt.Printf("Cursor")

	var result []presenter.Article

	rows, err := database.DB.Query(
		"SELECT id, title, description, content FROM articles LIMIT 10",
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		article := presenter.Article{}
		if err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content); err != nil {
			return nil, err
		}
		result = append(result, article)
	}

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) UpdateArticle(item *entities.Article) (*entities.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteArticle(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
