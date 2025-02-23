package users

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// User represents a user in the system
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FullName  string             `bson:"full_name" json:"full_name,omitempty"`
    Email     string             `bson:"email" json:"email,omitempty"`
    Password  string             `bson:"password,omitempty"` // Hashed password
    CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
