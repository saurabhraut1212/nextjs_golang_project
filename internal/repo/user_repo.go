package repo

import (
	"context"
	"errors"
	"time"

	"github.com/saurabhraut1212/nextjs_golang_project/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepo provides methods to interact with the user collection in MongoDB
type UserRepo struct{ col *mongo.Collection } // UserRepo is a struct that holds the MongoDB collection for users

// NewUserRepo creates a new UserRepo instance
func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		col: db.Collection("users"), // "users" is the name of the collection in MongoDB
	}
}

// find user by email
func (r *UserRepo) ByEmail(ctx context.Context, email string) (*models.User, error) { //ctx for timeout/cancellation
	var u models.User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &u, err
}

// find by id
func (r *UserRepo) ByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var u models.User
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &u, err
}

//create user

func (r *UserRepo) Create(ctx context.Context, u *models.User) error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = u.CreatedAt
	_, err := r.col.InsertOne(ctx, u)
	if mongo.IsDuplicateKeyError(err) {
		return errors.New("email already exists")
	}
	return err
}

func (r *UserRepo) EnsureIndexes(ctx context.Context) error {
	_, err := r.col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: optionsBoolUnique(true),
	})
	return err
}

func optionsBoolUnique(b bool) *options.IndexOptions {
	return options.Index().SetUnique(b)
}
