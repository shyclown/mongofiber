package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Item Constructs your Item model under entities.
type Item struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Items
type DeleteRequest struct {
	ID string `json:"id"`
}
