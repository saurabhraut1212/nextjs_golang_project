package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo represents a task in the todo list
type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"` // unique identifier, can be empty when creating a new todo
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`     // ID of the user who owns this todo
	Title     string             `bson:"title" json:"title"`
	Completed bool               `bson:"completed" json:"completed"` // indicates if the todo is completed
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"` // timestamp when the todo was created
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
