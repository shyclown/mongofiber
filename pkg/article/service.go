package article

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
)

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

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertArticle(item *entities.Article) (*entities.Article, error) {
	return s.repository.InsertArticle(item)
}
func (s *service) FetchArticles() (*[]presenter.Article, error) {

	return s.repository.FetchArticles()
}

func (s *service) FetchArticle(ID uuid.UUID) (*entities.Article, error) {
	return s.repository.FetchArticle(ID)
}

func (s *service) UpdateArticle(item *entities.Article) (*entities.Article, error) {
	return s.repository.UpdateArticle(item)
}
func (s *service) RemoveArticle(ID uuid.UUID) error {
	return s.repository.DeleteArticle(ID)
}
