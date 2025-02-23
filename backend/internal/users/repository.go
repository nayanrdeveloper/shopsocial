package users

import (
    "context"
    "errors"
    "shopsocial-backend/config"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
    Collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        Collection: config.GetCollection("users"),
    }
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *User) (*User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user.ID = primitive.NewObjectID()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    _, err := r.Collection.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// FindByEmail returns a user by email
func (r *UserRepository) FindByEmail(email string) (*User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user User
    err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}
