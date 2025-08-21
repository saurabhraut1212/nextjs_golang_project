// this file defines the user schema for mongodb database
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
// bson - how data is stored in mongodb
// json - how data is sent to the client
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"` //unique identifier and it can be empty while creating a new user
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password" json:"password"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
