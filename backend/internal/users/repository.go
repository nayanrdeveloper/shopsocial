package users

import (
	"context"
	"shopsocial-backend/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Collection: config.GetCollection("users"),
	}
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}
