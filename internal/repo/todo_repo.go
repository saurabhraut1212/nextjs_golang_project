package repo

import (
	"context"
	"time"

	"github.com/saurabhraut1212/nextjs_golang_project/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepo struct{ col *mongo.Collection }

func NewTodoRepo(db *mongo.Database) *TodoRepo { return &TodoRepo{col: db.Collection("todos")} }

func (r *TodoRepo) ListByUser(ctx context.Context, uid primitive.ObjectID) ([]models.Todo, error) {
	cur, err := r.col.Find(ctx, bson.M{"userId": uid}, options.Find().SetSort(bson.M{"createdAt": -1}))
	if err != nil {
		return nil, err
	}
	var out []models.Todo
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *TodoRepo) Create(ctx context.Context, uid primitive.ObjectID, title string) (models.Todo, error) {
	t := models.Todo{
		ID:        primitive.NewObjectID(),
		UserID:    uid,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	_, err := r.col.InsertOne(ctx, t)
	return t, err
}
