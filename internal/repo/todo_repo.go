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

func (r *TodoRepo) EnsureIndexes(ctx context.Context) error {
	// index: userId + createdAt (for fast listing)
	_, err := r.col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"userId": 1}},
		{Keys: bson.M{"createdAt": -1}},
	})
	return err
}

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

func (r *TodoRepo) Update(ctx context.Context, uid, id primitive.ObjectID, set bson.M) (models.Todo, error) {
	set["updatedAt"] = time.Now().UTC()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := r.col.FindOneAndUpdate(ctx, bson.M{"_id": id, "userId": uid}, bson.M{"$set": set}, opts)
	var out models.Todo
	err := res.Decode(&out)
	return out, err
}

func (r *TodoRepo) Delete(ctx context.Context, uid, id primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id, "userId": uid})
	return err
}
