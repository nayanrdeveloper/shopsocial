package products

import (
	"context"
	"errors"
	"shopsocial-backend/config"
	"shopsocial-backend/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		Collection: config.GetCollection("products"),
	}
}

func (r *ProductRepository) CreateProduct(product *Product) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := r.Collection.InsertOne(ctx, product)
	if err != nil {
		logger.Log.Error("Failed to insert product", zap.Error(err))
		return nil, err
	}

	logger.Log.Info("Product inserted successfully", zap.String("id", product.ID.Hex()))
	return product, nil
}

func (r *ProductRepository) GetProductByID(id string) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Warn("Invalid product ID format", zap.String("id", id))
		return nil, errors.New("invalid product ID format")
	}

	var product Product
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		logger.Log.Warn("Product not found", zap.String("id", id))
		return nil, err
	}

	logger.Log.Info("Product retrieved successfully", zap.String("id", id))
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(id string, updateData bson.M) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Warn("Invalid product ID format", zap.String("id", id))
		return nil, errors.New("invalid product ID format")
	}

	updateData["updated_at"] = time.Now()
	update := bson.M{"$set": updateData}

	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		logger.Log.Error("Failed to update product", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	logger.Log.Info("Product updated successfully", zap.String("id", id))
	return r.GetProductByID(id)
}

func (r *ProductRepository) DeleteProduct(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Warn("Invalid product ID format", zap.String("id", id))
		return errors.New("invalid product ID format")
	}

	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		logger.Log.Error("Failed to delete product", zap.String("id", id), zap.Error(err))
		return err
	}

	logger.Log.Info("Product deleted successfully", zap.String("id", id))
	return nil
}
