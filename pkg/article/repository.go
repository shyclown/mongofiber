package article

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/database"
	"mongofiber/pkg/entities"
)

type Repository interface {
	InsertArticle(article *entities.Article) (*entities.Article, error)
	FetchArticles() (*[]presenter.Article, error)
	FetchArticle(ID uuid.UUID) (*entities.Article, error)
	UpdateArticle(article *entities.Article) (*entities.Article, error)
	DeleteArticle(ID uuid.UUID) error
}

type repository struct {
	Table string
}

// NewRepo is the single instance repo that is being created.
func NewRepo(table string) Repository {
	return &repository{
		Table: table,
	}
}

func (r *repository) InsertArticle(article *entities.Article) (*entities.Article, error) {
	article.ID = uuid.New()
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

	rows, err := database.DB.Query(
		"SELECT id, title, description, content FROM articles LIMIT 10",
	)
	if err != nil {
		return nil, err
	}
	var result []presenter.Article
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

func (r *repository) FetchArticle(ID uuid.UUID) (*entities.Article, error) {
	rows, err := database.DB.Query(
		"SELECT id, title, description, content FROM articles WHERE id=?", ID.String(),
	)
	if err != nil {
		return nil, err
	}
	var result []entities.Article
	for rows.Next() {
		article := entities.Article{}
		if err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content); err != nil {
			return nil, err
		}
		result = append(result, article)
	}
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func (r *repository) UpdateArticle(article *entities.Article) (*entities.Article, error) {
	_, err := database.DB.Query(
		"UPDATE articles SET title = ?, description = ?, content= ? WHERE id = ?",
		article.Title, article.Description, article.Content, article.ID,
	)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *repository) DeleteArticle(ID uuid.UUID) error {
	_, err := database.DB.Query(
		"DELETE FROM articles WHERE id=?", ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}
