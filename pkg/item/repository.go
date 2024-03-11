package item

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongofiber/api/presenter"
	"mongofiber/pkg/entities"
	"time"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateItem(item *entities.Item) (*entities.Item, error)
	ReadItem() (*[]presenter.Item, error)
	UpdateItem(item *entities.Item) (*entities.Item, error)
	DeleteItem(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// CreateItem is a mongo repository that helps to create items
func (r *repository) CreateItem(item *entities.Item) (*entities.Item, error) {
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// ReadItem is a mongo repository that helps to fetch items
func (r *repository) ReadItem() (*[]presenter.Item, error) {
	var items []presenter.Item
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var item presenter.Item
		_ = cursor.Decode(&item)
		items = append(items, item)
	}
	return &items, nil
}

// UpdateItem is a mongo repository that helps to update items
func (r *repository) UpdateItem(item *entities.Item) (*entities.Item, error) {
	item.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": item.ID}, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteItem is a mongo repository that helps to delete items
func (r *repository) DeleteItem(ID string) error {
	itemID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": itemID})
	if err != nil {
		return err
	}
	return nil
}
