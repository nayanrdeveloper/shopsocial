package products

import (
	"net/http"
	"shopsocial-backend/pkg/logger"
	"shopsocial-backend/pkg/responses"

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

// CREATE
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		responses.SendError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	createdProduct, err := h.Service.CreateProduct(&product)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	logger.Log.Info("Product created successfully", zap.String("id", createdProduct.ID.Hex()))
	responses.SendCreated(c, "Product created successfully", createdProduct)
}

// READ
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Fetching product", zap.String("id", id))

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		responses.SendError(c, http.StatusNotFound, "Product not found", err)
		return
	}

	responses.SendSuccess(c, "Product retrieved successfully", product)
}

// UPDATE
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateData bson.M
	if err := c.ShouldBindJSON(&updateData); err != nil {
		responses.SendError(c, http.StatusBadRequest, "Invalid update data", err)
		return
	}

	updatedProduct, err := h.Service.UpdateProduct(id, updateData)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	logger.Log.Info("Product updated successfully", zap.String("id", id))
	responses.SendSuccess(c, "Product updated successfully", updatedProduct)
}

// DELETE
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Deleting product", zap.String("id", id))

	err := h.Service.DeleteProduct(id)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	logger.Log.Info("Product deleted successfully", zap.String("id", id))
	responses.SendDeleted(c, "Product deleted successfully")
}
