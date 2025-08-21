package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(mongoURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //creates a context that automatically cancels after 10 seconds
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
}
