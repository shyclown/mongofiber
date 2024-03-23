package item

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
)

type Service interface {
	InsertItem(item *entities.Item) (*entities.Item, error)
	FetchItems() (*[]presenter.Item, error)
	FetchItem(ID uuid.UUID) (*entities.Item, error)
	UpdateItem(item *entities.Item) (*entities.Item, error)
	RemoveItem(ID uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertItem(item *entities.Item) (*entities.Item, error) {
	return s.repository.CreateItem(item)
}
func (s *service) FetchItems() (*[]presenter.Item, error) {
	return s.repository.ReadItems()
}
func (s *service) FetchItem(ID uuid.UUID) (*entities.Item, error) {
	return s.repository.ReadItem(ID)
}
func (s *service) UpdateItem(item *entities.Item) (*entities.Item, error) {
	return s.repository.UpdateItem(item)
}
func (s *service) RemoveItem(ID uuid.UUID) error {
	return s.repository.DeleteItem(ID)
}
