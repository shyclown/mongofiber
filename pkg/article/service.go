package article

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
)

type Service interface {
	CreateArticle(article *entities.Article) (*entities.Article, error)
	ReadArticles() (*[]presenter.Article, error)
	ReadArticle(ID uuid.UUID) (*entities.Article, error)
	UpdateArticle(article *entities.Article) (*entities.Article, error)
	DeleteArticle(ID uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateArticle(article *entities.Article) (*entities.Article, error) {
	return s.repository.InsertArticle(article)
}

func (s *service) ReadArticles() (*[]presenter.Article, error) {

	return s.repository.FetchArticles()
}

func (s *service) ReadArticle(ID uuid.UUID) (*entities.Article, error) {
	return s.repository.FetchArticle(ID)
}

func (s *service) UpdateArticle(item *entities.Article) (*entities.Article, error) {
	return s.repository.UpdateArticle(item)
}
func (s *service) DeleteArticle(ID uuid.UUID) error {
	return s.repository.DeleteArticle(ID)
}
