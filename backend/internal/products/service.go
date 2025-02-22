package products

import (
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductService struct {
	Repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(product *Product) (*Product, error) {
	if product.Price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}
	if product.Stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id string) (*Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(id string, updateData bson.M) (*Product, error) {
	return s.Repo.UpdateProduct(id, updateData)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.DeleteProduct(id)
}
