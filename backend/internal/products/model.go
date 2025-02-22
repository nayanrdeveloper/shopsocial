package products

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product represents a product entity
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price" validate:"required,gt=0"`
	Category    string             `bson:"category" json:"category" validate:"required"`
	Stock       int                `bson:"stock" json:"stock" validate:"gte=0"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
