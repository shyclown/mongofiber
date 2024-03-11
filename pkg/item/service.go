package item

import (
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertItem(item *entities.Item) (*entities.Item, error)
	FetchItems() (*[]presenter.Item, error)
	UpdateItem(item *entities.Item) (*entities.Item, error)
	RemoveItem(ID string) error
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

// InsertItem is a service layer that helps insert item in ItemShop
func (s *service) InsertItem(item *entities.Item) (*entities.Item, error) {
	return s.repository.CreateItem(item)
}

// FetchItems is a service layer that helps fetch all items in ItemShop
func (s *service) FetchItems() (*[]presenter.Item, error) {
	return s.repository.ReadItem()
}

// UpdateItem is a service layer that helps update items in ItemShop
func (s *service) UpdateItem(item *entities.Item) (*entities.Item, error) {
	return s.repository.UpdateItem(item)
}

// RemoveItem is a service layer that helps remove items from ItemShop
func (s *service) RemoveItem(ID string) error {
	return s.repository.DeleteItem(ID)
}
