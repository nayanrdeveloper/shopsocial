package products

import (
	"net/http"
	"shopsocial-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		logger.Log.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.Service.CreateProduct(&product)
	if err != nil {
		logger.Log.Error("Failed to create product", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info("Product created successfully", zap.String("id", createdProduct.ID.Hex()))
	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Fetching product", zap.String("id", id))

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		logger.Log.Warn("Product not found", zap.String("id", id))
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	logger.Log.Info("Product retrieved successfully", zap.String("id", id))
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateData bson.M
	if err := c.ShouldBindJSON(&updateData); err != nil {
		logger.Log.Warn("Invalid update data", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.Service.UpdateProduct(id, updateData)
	if err != nil {
		logger.Log.Error("Failed to update product", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info("Product updated successfully", zap.String("id", id))
	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Deleting product", zap.String("id", id))

	err := h.Service.DeleteProduct(id)
	if err != nil {
		logger.Log.Error("Failed to delete product", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info("Product deleted successfully", zap.String("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
