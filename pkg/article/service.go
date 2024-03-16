package article

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertArticle(article *entities.Article) (*entities.Article, error)
	FetchArticles() (*[]presenter.Article, error)
	FetchArticle(ID uuid.UUID) (*entities.Article, error)
	UpdateArticle(article *entities.Article) (*entities.Article, error)
	RemoveArticle(ID uuid.UUID) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertArticle is a service layer that helps insert item in ArticleShop
func (s *service) InsertArticle(item *entities.Article) (*entities.Article, error) {
	return s.repository.InsertArticle(item)
}

// FetchArticles is a service layer that helps fetch all articles
func (s *service) FetchArticles() (*[]presenter.Article, error) {
	return s.repository.FetchArticles()
}

// FetchArticle is a service layer that helps fetch single article
func (s *service) FetchArticle(ID uuid.UUID) (*entities.Article, error) {
	return s.repository.FetchArticle(ID)
}

// UpdateArticle is a service layer that helps update items in ArticleShop
func (s *service) UpdateArticle(item *entities.Article) (*entities.Article, error) {
	return s.repository.UpdateArticle(item)
}

// RemoveArticle is a service layer that helps remove items from ArticleShop
func (s *service) RemoveArticle(ID uuid.UUID) error {
	return s.repository.DeleteArticle(ID)
}
