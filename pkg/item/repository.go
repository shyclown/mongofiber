package item

import (
	"github.com/google/uuid"
	"mongofiber/api/presenter"
	"mongofiber/database"
	"mongofiber/pkg/entities"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateItem(item *entities.Item) (*entities.Item, error)
	ReadItems() (*[]presenter.Item, error)
	ReadItem(ID uuid.UUID) (*entities.Item, error)
	UpdateItem(item *entities.Item) (*entities.Item, error)
	DeleteItem(ID uuid.UUID) error
	GetItemByEntityId(ID uuid.UUID) (*entities.Item, error)
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

// CreateItem is a mongo repository that helps to create items
func (r *repository) CreateItem(item *entities.Item) (*entities.Item, error) {
	item.Id = uuid.New()
	_, err := database.DB.Query(
		"INSERT INTO items (id, title, description, entity_id, entity_type) VALUES (?, ?, ?, ?, ?)",
		item.Id, item.Title, item.Description, item.EntityId, item.EntityType,
	)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) ReadItems() (*[]presenter.Item, error) {

	rows, err := database.DB.Query(
		"SELECT id, title, description, entity_id, entity_type FROM items LIMIT 10",
	)
	if err != nil {
		return nil, err
	}
	var result []presenter.Item
	for rows.Next() {
		item := presenter.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.EntityId, &item.EntityType); err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	if err != nil {

		return nil, err
	}
	return &result, nil
}

func (r *repository) ReadItem(ID uuid.UUID) (*entities.Item, error) {
	rows, err := database.DB.Query(
		"SELECT id, title, description, entity_id, entity_type FROM items WHERE id=?", ID.String(),
	)
	if err != nil {
		return nil, err
	}
	var result []entities.Item
	for rows.Next() {
		item := entities.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.EntityId, &item.EntityType); err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func (r *repository) UpdateItem(item *entities.Item) (*entities.Item, error) {
	_, err := database.DB.Query(
		"UPDATE items SET title = ?, description = ? WHERE id = ?",
		item.Title, item.Description, item.Id,
	)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) DeleteItem(ID uuid.UUID) error {
	_, err := database.DB.Query(
		"DELETE FROM items WHERE id=?",
		ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetItemByEntityId(ID uuid.UUID) (*entities.Item, error) {
	rows, err := database.DB.Query(
		"SELECT id, title, description, entity_id, entity_type FROM items WHERE entity_id=?", ID.String(),
	)
	if err != nil {
		return nil, err
	}
	var result []entities.Item
	for rows.Next() {
		item := entities.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.EntityId, &item.EntityType); err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}
